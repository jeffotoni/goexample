package main

import (
    "bytes"
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "sync"
    "time"

    g "github.com/jeffotoni/gconcat"
    "github.com/lib/pq"
)

var (
    dbLocal *sql.DB
    once    sync.Once
    err     error
)

var (
    DB_NAME     = os.Getenv("DB_NAME")
    DB_HOST     = os.Getenv("DB_HOST")
    DB_USER     = os.Getenv("DB_USER")
    DB_PASSWORD = os.Getenv("DB_PASSWORD")
    DB_PORT     = os.Getenv("DB_PORT")

    DB_SSL       = "require"
    DB_CERT_PATH = ""
    DB_SORCE     = "postgres"
)

type data struct {
    Key       int
    Timestamp string
    Name      string
    Status    int
}

type events struct {
    Table  string
    Action string
    Data   data
}

var e events

func ConnectNew() *sql.DB {
    once.Do(func() {
        if dbLocal != nil {
            return
        }
        var DBINFO string
        DB_SSL = "require"
        DBINFO = g.ConcatStr("host=", DB_HOST, " port=", DB_PORT, " user=", DB_USER,
            " password=", DB_PASSWORD, " dbname=", DB_NAME,
            " sslmode=", DB_SSL)
        if dbLocal, err = sql.Open(DB_SORCE, DBINFO); err != nil {
            log.Println(err.Error())
        }
    })
    return dbLocal
}

func waitForNotification(db *sql.DB, l *pq.Listener) {
    for {
        select {
        case nj := <-l.Notify:
            fmt.Println("Recive: ", nj)

            var prettyJSON bytes.Buffer
            err := json.Indent(&prettyJSON, []byte(nj.Extra), "", "\t")
            if err != nil {
                fmt.Println("Error processing JSON: ", err)
                break
            }
            fmt.Println(string(prettyJSON.Bytes()))

            err = json.Unmarshal([]byte(nj.Extra), &e)
            if err != nil {
                fmt.Println("Error Unmarshal JSON: ", err)
                break
            }
            fmt.Println("----------------------------> ", e.Table)
            fmt.Println("----------------------------> ", e.Action)
            fmt.Println("----------------------------> ", e.Data.Key)
            fmt.Println("----------------------------> ", e.Data.Timestamp)
            fmt.Println("----------------------------> ", e.Data.Name)
            println("Enviando para Kafka...")
            time.Sleep(time.Second * 6)
            _, err = db.Exec("delete from public.expire_table WHERE key=$1", e.Data.Key)
            fmt.Println("del: ", err)
            println("removendo da fila postgres..")
            time.Sleep(time.Second * 2)
            break

        case <-time.After(10 * time.Second):
            go l.Ping()
            // Check if there's more work available, just in case it takes
            // a while for the Listener to notice connection loss and
            // reconnect.
            fmt.Println("received no work for 5 seconds, checking for new work")
        }
    }
}

func main() {
    Db := ConnectNew()
    var conninfo string = g.ConcatStr("user=", DB_NAME, "password=", DB_PASSWORD)
    reportProblem := func(ev pq.ListenerEventType, err error) {
        if err != nil {
            fmt.Println("ListenEvent: ", err.Error())
        }
    }

    minReconn := 10 * time.Second
    maxReconn := time.Minute
    listener := pq.NewListener(conninfo, minReconn, maxReconn, reportProblem)
    err := listener.Listen("events")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Loop Listen Postgresql")
    for {
        // process all available work before waiting for notifications
        waitForNotification(Db, listener)
    }
}
