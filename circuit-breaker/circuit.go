package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/cep21/circuit/closers/hystrix"
	"github.com/cep21/circuit/v3"
)

func main() {
	start := time.Now()
	go func() {
		time.Sleep(time.Second * 5)
		client := NewUserClient()
		for {
			time.Sleep(time.Millisecond * 100)
			usr, err := client.Get("http://localhost:8090/user")
			if err != nil {
				log.Println("Err to Get User: ", err.Error())
				continue
			}

			log.Printf("Response: %#v", usr)
		}
	}()

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		if (int(time.Now().Unix())-int(start.Unix()))%20 > 10 {
			w.WriteHeader(500)
			w.Write([]byte("internal server error"))
			return
		}

		w.Header().Add("Content-Type", "application/json")
		user := User{
			Name:  generateRandomString(7),
			Email: generateRandomString(10) + "@" + generateRandomString(4) + ".com",
		}

		b, err := json.Marshal(user)
		if err != nil {
			w.Write([]byte("internal server error"))
			w.WriteHeader(500)
			return
		}

		w.Write(b)
	})

	http.ListenAndServe("0.0.0.0:8090", nil)
}

type userClient struct {
	circuitBreaker *circuit.Circuit
	client         *http.Client
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserClient() *userClient {
	manager := circuit.Manager{}
	return &userClient{
		circuitBreaker: manager.MustCreateCircuit("user", circuit.Config{
			General: circuit.GeneralConfig{
				ClosedToOpenFactory: hystrix.OpenerFactory(hystrix.ConfigureOpener{
					ErrorThresholdPercentage: 20,
				}),
				OpenToClosedFactory: hystrix.CloserFactory(hystrix.ConfigureCloser{SleepWindow: time.Second * 2}),
			},
		}),
		client: http.DefaultClient,
	}
}

func (uc *userClient) Get(url string) (*User, error) {
	var user User
	err := uc.circuitBreaker.Execute(context.Background(), func(ctx context.Context) error {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return err
		}

		res, err := uc.client.Do(req)
		if err != nil {
			return err
		}

		if res.StatusCode != 200 {
			return errors.New(fmt.Sprint("err status: ", res.StatusCode))
		}

		b, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(b, &user)
		if err != nil {
			return err
		}

		return nil

	}, uc.fallback)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (uc *userClient) fallback(ctx context.Context, err error) error {
	if uc.circuitBreaker.IsOpen() {
		return errors.New("circuit is open")
	}

	return err
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(b)
}
