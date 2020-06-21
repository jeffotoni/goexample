package pub

import (
	"context"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/pubsub"
)

var (
	projectID = "projeto-eng1"
	nameTopic = "topic.pub.go"
)

func Pub(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "Error Decode..")
		return
	}
	if d.Name == "" {
		w.WriteHeader(400)
		fmt.Fprint(w, "Error Name é obrigatorio!")
		return
	}

	id, err := publish(html.EscapeString(d.Name))
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "Error:", err.Error())
		return
	}

	w.WriteHeader(200)
	fmt.Fprintf(w, "Enviado com sucesso pub ID=[%s]!", id)
}

func publish(msg string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("Error NewClient:%s", err.Error())
		return "", err
	}

	topic := client.Topic(nameTopic)
	defer topic.Stop()

	res := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(msg)},
	)

	id, err := res.Get(ctx)
	if err != nil {
		log.Printf("Error Get results:%s", err.Error())
		return "", err
	}

	return string(id), nil
}
