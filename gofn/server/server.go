package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

///
import (
	"database/sql"
	_ "github.com/lib/pq"
)

var once sync.Once

var pchan = make(chan string)

const LayoutDateLog = "2006-01-02 15:04:05"
const LayoutDate = "2006-01-02"
const LayoutHour = "15:04:05"

const (
	DB_HOST     = "192.168.35.251"
	DB_USER     = "gofn"
	DB_PASSWORD = "1234"

	DB_PORT  = "5432"
	DB_SSL   = "disable"
	DB_SORCE = "postgres"
)

type PgStruct struct {
	Pgdb *sql.DB
}

type User struct {
	Id    string `json:id`
	Name  string `json:name`
	Email string `json:email`
	Db    string `json:db`
	Msg   string `json:msg`
}

// cache sync.Map
type cache struct {
	mm sync.Map
	sync.Mutex
}

var (
	err         error
	PostDb      PgStruct
	PORT_SERVER = "9010" // definir no init
)

var (
	pool = &cache{}
)

// put sync.Map
func (c *cache) put(key, value interface{}) {

	c.Lock()
	defer c.Unlock()
	c.mm.Store(key, value)
}

// get sync.Map
func (c *cache) get(key interface{}) interface{} {

	c.Lock()
	defer c.Unlock()

	v, _ := c.mm.Load(key)
	return v

}

// setLoad... fn func() interface{}
func (c *cache) loadStore(key interface{}, fc func() interface{}) interface{} {

	c.Lock()
	defer c.Unlock()

	if v, ok := c.mm.Load(key); ok {
		return v
	}

	val := fc()
	c.mm.Store(key, val)
	return val
}

// conectando de forma segura usando goroutine
func PgConnect(DBNAME, DB_HOST, DB_USER, DB_PASSWORD string) interface{} {

	// existe o banco
	if dbPg := pool.get(DBNAME); dbPg != nil {

		// return objeto conexao
		return dbPg.(*sql.DB)

	} else {

		// removendo aspas..
		DBNAME = strings.Replace(DBNAME, `"`, "", -1)

		DBINFO := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DBNAME, DB_SSL)

		// log.Println(DBINFO)

		// func para ser executada
		// dentro do loadStore
		// quando duas ou mais
		// goroutine chegarem
		// neste mesmo momento
		// de fazer um Store
		fn := func() interface{} {

			PostDb.Pgdb, err = sql.Open(DB_SORCE, DBINFO)

			if err != nil {

				errordb := fmt.Sprintf("Unable to connection to database: %v\n", err)
				log.Println("error:: ", errordb)
				defer PostDb.Pgdb.Close()
				return nil
			}

			if ok2 := PostDb.Pgdb.Ping(); ok2 != nil {

				log.Println("connect error...: ", ok2)
				defer PostDb.Pgdb.Close()
				return nil
			}

			log.Println("connect return sucess:: client [" + DBNAME + "]")
			return PostDb.Pgdb
		}

		// recebendo conexao

		// armazenando cache loadStore
		sqlDb := pool.loadStore(DBNAME, fn)

		if sqlDb != nil {

			return sqlDb.(*sql.DB)

		} else {

			return nil
		}
	}
}

func (dbx *PgStruct) PgPing() error {

	db := dbx.Pgdb

	if err := db.Ping(); err == nil {

		return nil

	} else {

		return err
	}
}

func CacheUpdate(tipo, IdUserDb, Dbname string, p chan<- string) {

	var prepare string

	if tipo == "success" {

		prepare = "update login set sucess=sucess+1 where id=$1"

	} else {

		prepare = "update login set error=error+1 where id=$1"
	}

	prepare = prepare + "#" + IdUserDb + "#" + Dbname

	p <- prepare
}

func goPgUpdate(prepare string) {

	// aguardando um pouco
	// para consumir a fila
	time.Sleep(time.Millisecond * 10)

	var errx error
	var update *sql.Stmt

	Vet := strings.Split(prepare, "#")

	// update
	UpdateString := Vet[0]

	ID_USER, errc := strconv.Atoi(Vet[1]) // strint to int

	if errc != nil {

		log.Println("Falta o Id do User...", errc)
		return
	}

	DbName := decbase64(Vet[2]) // Nome do Banco encode64

	// local
	var Db *sql.DB

	// Db...
	if interf := PgConnect(DbName, DB_HOST, DB_USER, DB_PASSWORD); interf != nil {

		Db = interf.(*sql.DB)

	} else {

		log.Println("error ao fazer connect goPgUpdate com Db..", interf)
		return
	}

	update, errx = Db.Prepare(UpdateString)

	if errx != nil {

		log.Println("erro ao fazer prepare update: ", errx)
		return
	}

	defer update.Close() // danger!

	res, err := update.Exec(ID_USER)

	if err != nil {

		log.Printf("Error update.Exec  user %v\n", err)
		return

	} else {

		afetUp, err := res.RowsAffected()

		if err != nil {

			log.Println("Error update.Exec Affected rows user:::", err)
		}

		// fmt.Println(afetUp)
		log.Println("update success:: ", fmt.Sprintf("%d", afetUp), " ::iduser:: ", ID_USER, " :: Database :: ", DbName)
	}
}

func Now() string {

	Now := time.Now().Format(LayoutDateLog)
	Now = fmt.Sprintf("%v", Now)

	return Now
}

func DateHora() string {

	return fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04:05"))
}

func Date() string {

	return fmt.Sprintf("%s", time.Now().Format("2006-01-02"))
}

func goApiUpdate(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") != "application/json" {

		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Error content-type")
		return
	}

	if r.Header.Get("X-key") != "1234567890x1234#4" {

		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Error de chave em seu json!")
		return
	}

	// Bearer ac2168444f4de69c27d6384ea2ccf61a49669be5a2fb037ccc1f
	if r.Header.Get("Authorization") != "Bearer ac2168444f4de69c27d6384ea2ccf61a49669be5a2fb037ccc1f" {

		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Error de Authorization em seu json!")
		return
	}

	if strings.ToUpper(r.Method) == "POST" {

		b, err := ioutil.ReadAll(r.Body)

		defer r.Body.Close()

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		//vetMap := make(map[string]interface{})
		var v = &User{}

		err = json.Unmarshal(b, &v)

		if err != nil {

			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, "Error ao fazer merge em struct")
			return
		}

		// removendo aspas..
		v.Db = strings.Replace(v.Db, `"`, "", -1)

		if strings.ToUpper(v.Msg) == "OK" {

			//PgUpdate("success", r.RemoteAddr)

			go CacheUpdate("success", v.Id, v.Db, pchan)
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "update feito com sucess: "+v.Name) // send data to client side

		} else {

			//PgUpdate("error", r.RemoteAddr)
			go CacheUpdate("error", v.Id, v.Db, pchan)
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, "update a ser feito em error: "+v.Name) // send data to client side
		}

	} else {

		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Method tem que ser POST..") // send data to client side
	}
}

func maxClientsFunc(h http.Handler, n int) http.Handler {

	sema := make(chan struct{}, n)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sema <- struct{}{}

		defer func() { <-sema }()

		h.ServeHTTP(w, r)
	})
}

// decode base64
func decbase64(str64 string) string {

	data, err := base64.StdEncoding.DecodeString(str64)
	if err != nil {
		log.Println("error base64:", err, " Dado: ", str64)
		return str64
	}

	return fmt.Sprintf("%q", data)
}

func init() {

	////////// inicio
	port_tmp := os.Getenv("PORT_SERVER")

	if port_tmp != "" {

		PORT_SERVER = port_tmp

	} else {

		//if for argumentos OK
		if len(os.Args) == 2 && os.Args[1] != "" {

			PORT_SERVER = os.Args[1]

		}
	}

	log.Println("server run port: " + PORT_SERVER) // print

	////////////////////////////////
	///
}

// api-server
// @jeffotoni
func main() {

	const maxClients = 20000

	// handler /v1/api/ses
	handlerApiUp := http.HandlerFunc(goApiUpdate)

	// fazendo o controle de conexoes
	http.Handle("/api/user", maxClientsFunc(handlerApiUp, maxClients))

	srv := &http.Server{

		Addr: ":" + PORT_SERVER,
		// timeout para o nosso http
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		err := srv.ListenAndServe() // set listen port
		if err != nil {
			log.Println("ListenAndServe: ", err)
		}
	}()

	// consumir sempre que
	// tiver no channel
	// Uma fila de Update
	for ppc := range pchan {

		goPgUpdate(ppc)
	}

	defer close(pchan)
}
