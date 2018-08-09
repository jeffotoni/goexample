package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	//"sync"
	"time"
)

import (
	"github.com/boltdb/bolt"
	//"github.com/jeffotoni/s3fileworkclient/pkg/gbolt"
	"github.com/jeffotoni/s3fileworkclient/pkg/config"
	//"log"
	//"time"
	"github.com/jeffotoni/s3fileworkclient/pkg/timestamp"
)

type LoginData struct {
	Key     string `json:"key"`
	Login   string `json:"login"`
	Created string `json:"created"`
	Success bool   `json:"success"`
	// quantas vezes logou
	Count int `json:"count"`
}

var LoginStruct LoginData

// config do noSql
// ira determinar o momento de fazer as varreduras ou nao
// teremos um time para isto ocorrer
type JsonDataDbConf struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	Created string `json:"created"`

	// quantas vezes abriu
	Count int `json:"count"`
}

// We created a json date type of our structure
var jDbConf JsonDataDbConf

func main() {

	// db, err := bolt.Open("./db/gbolt.db", 0666, nil)
	// if err != nil {
	// 	log.Println("Error: ", err)
	// }

	//defer db.Close()

	//db := gbolt.Connect(true)

	//fmt.Println("Vamos ver: ping:::  ", gbolt.Ping())
	fmt.Println("Salvando user")
	//gbolt.SaveLogin("jeff.otoni@gmail.com", false)
	SaveLogin("jeff.otoni@gmail.com", false)
}

func GetStringJson(keyS string) ([]byte, error) {

	var byteVal []byte
	//var stringVal string

	reset := make(chan string, 1)
	reset2 := make(chan []byte, 1)

	//db := Connect(ReadOnly)
	//Tem que ser individual por metodo
	//cada metodo em seu open, para permitir outras aplicoes acessarem o banco
	db, err := bolt.Open(os.Getenv("PATH_BOLTDB")+config.PATH_DB, 0600, &bolt.Options{Timeout: 5 * time.Second})

	if err != nil {

		return byteVal, errors.New("error noSql Open boltDB")
	}

	defer db.Close()

	key := []byte(keyS)

	//fmt.Println(key)

	//var mutex *sync.Mutex
	//var mutex = &sync.Mutex{}

	err = db.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket(config.DATABASE)

		if bucket == nil {

			return errors.New("Bucket not found!")
		}

		xbyte := bucket.Get(key)

		//mutex.Lock()
		//byteVal = xbyte
		//stringVal = string(xbyte)
		reset <- string(xbyte)

		reset2 <- xbyte

		//fmt.Println(stringVal)

		//byteVal = xbyte
		//defer mutex.Unlock()

		return nil
	})

	stringVal := <-reset

	if err != nil {

		return byteVal, errors.New("error noSql Get not found!")
	}

	if stringVal == "" {

		return byteVal, errors.New("error, no existe conteudo no noSql ")
	}

	//fmt.Sprintf("%v", byteVal)

	// liberar
	// defer db.Close()

	//byt := []byte(string(valbyte))
	/// interface
	//errjs := json.Unmarshal(byt, &djson)
	//if errjs != nil {

	//fmt.Println("error json: ", errjs)
	//}

	byteNew := []byte(stringVal)

	byteVal2 := <-reset2

	//fmt.Println("vamos ver:..", valbyte)
	fmt.Println("string:::", stringVal)
	fmt.Println("byte  ::: ", byteNew)
	fmt.Println("byte 2 chan  ::: ", byteVal2)
	//os.Exit(0)

	return byteNew, nil
}

// SaveDb This method prepares the whole jsonstring to save in boltdb
func SaveLogin(Login string, Success bool) error {

	keyfile := "login"

	count := 0

	now := fmt.Sprintf("%s", timestamp.FromNow{})

	dataObj := JsonObj(keyfile)

	if dataObj.Count > 0 {

		count = dataObj.Count + 1

	} else {

		count = 1
	}

	stringJson := LoginData{keyfile, Login, now, Success, count}

	respJson, err := json.Marshal(stringJson)

	respJsonX := string(respJson)

	err = Save(keyfile, respJsonX)

	if err == nil {

		return nil

	} else {

		return err
	}
}

// Save This method is responsible for saving on boltdb
func Save(keyS string, valueS string) error {

	//db := Connect(WriteOnly)
	//db := OpenWrite()

	db, err := bolt.Open(config.PATH_DB, 0600, &bolt.Options{Timeout: 5 * time.Second})

	if err != nil {

		return errors.New("error noSql Open boltDB")
	}

	defer db.Close()

	key := []byte(keyS)
	value := []byte(valueS)

	err = db.Update(func(tx *bolt.Tx) error {

		bucket, err := tx.CreateBucketIfNotExists(config.DATABASE)

		if err != nil {

			return err
		}

		err = bucket.Put(key, value)

		if err != nil {

			return err

		} else {

			return nil
		}
	})

	if err != nil {

		fmt.Println("Error try save in noSql: ", err)
		os.Exit(0)
	}

	// liberar
	//db.Close()

	return nil
}

// JsonGet This method is responsible for returning the
// content in json format]
func JsonObj(keyS string) *JsonDataDbConf {

	//var valbyte2 []byte
	jsonx := JsonDataDbConf{}

	valbyte, err := GetStringJson(keyS)

	fmt.Println("vamos save now.. [GetStringJson]", err)
	fmt.Println("vamos save now.. [GetStringJson]", valbyte)
	os.Exit(0)

	if err != nil {

		fmt.Println("Error no JsonObj: ", err)
		return &jsonx
	}

	// // convert object json
	json.Unmarshal(valbyte, &jsonx)

	//fmt.Println(jDbConf.Name)

	return &jsonx
}

/** var wg sync.WaitGroup

//for i := 0; i < 1; i++ {

wg.Add(1)

go func() {
	defer wg.Done()

	// Do some work in parallel...

	// *** This will block until the batch is finally run ***
	err = db.Batch(func(tx *bolt.Tx) error {

		err := db.View(func(tx *bolt.Tx) error {
			//var mutex = &sync.Mutex{}

			bucket := tx.Bucket(config.DATABASE)

			if bucket == nil {

				return errors.New("Bucket not found!")
			}

			xbyte := bucket.Get(key)

			fmt.Println("aqui::::", xbyte)
			fmt.Println("aqui::::", string(xbyte))

			//mutex.Lock()
			stringVal = string(xbyte)
			//defer mutex.Unlock()

			return nil
		})

		return err
	})

	if err != nil {
		fmt.Printf("Got an error: %s\n", err)
	}
}()
//}

wg.Wait()
db.Close()
*/
