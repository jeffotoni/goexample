package main_test

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gofiber/fiber"
)

const jsonExample = `{
    "glossary": {
        "title": "example glossary",
        "GlossDiv": {
            "title": "S",
            "GlossList": {
                "GlossEntry": {
                    "ID": "SGML",
                    "SortAs": "SGML",
                    "GlossTerm": "Standard Generalized Markup Language",
                    "Acronym": "SGML",
                    "Abbrev": "ISO 8879:1986",
                    "GlossDef": {
                        "para": "A meta-markup language, used to create markup languages such as DocBook.",
                        "GlossSeeAlso": ["GML", "XML"]
                    },
                    "GlossSee": "markup"
                }
            }
        }
    }
}`

type Glossary struct {
    Glossary struct {
        Title    string `json:"title"`
        GlossDiv struct {
            Title     string `json:"title"`
            GlossList struct {
                GlossEntry struct {
                    ID        string `json:"ID"`
                    SortAs    string `json:"SortAs"`
                    GlossTerm string `json:"GlossTerm"`
                    Acronym   string `json:"Acronym"`
                    Abbrev    string `json:"Abbrev"`
                    GlossDef  struct {
                        Para         string   `json:"para"`
                        GlossSeeAlso []string `json:"GlossSeeAlso"`
                    } `json:"GlossDef"`
                    GlossSee string `json:"GlossSee"`
                } `json:"GlossEntry"`
            } `json:"GlossList"`
        } `json:"GlossDiv"`
    } `json:"glossary"`
}

type Book struct {
    Title    string   `json:"title"`
    Author   string   `json:"author"`
    Pages    int      `json:"num_pages"`
    Chapters []string `json:"chapters"`
}

type BookDef struct {
    Title  string `msg:"title"`
    Author string `msg:"author"`
    Pages  int    `msg:"num_pages"`
}

func BenchmarkWriteJson(b *testing.B) {
    for n := 0; n < b.N; n++ {
        w := httptest.NewRecorder()
        var book = &Book{Title: "Titulo teste", Author: "jeffotoni", Pages: 200, Chapters: []string{"C", "C++", "Algol", "Basic"}}
        err := json.NewEncoder(w).Encode(book)
        if err != nil {
            log.Println(err)
        }
    }
}

func BenchmarkFprintfJsonMarshal(b *testing.B) {
    for n := 0; n < b.N; n++ {
        w := httptest.NewRecorder()
        var book = &Book{Title: "Titulo teste", Author: "jeffotoni", Pages: 200, Chapters: []string{"C", "C++", "Algol", "Basic"}}
        b, err := json.Marshal(book)
        if err != nil {
            log.Println(err)
            continue
        }
        fmt.Fprint(w, string(b))
    }
}

func BenchmarkWriteStringJsonMarshal(b *testing.B) {
    for n := 0; n < b.N; n++ {
        w := httptest.NewRecorder()
        var book = &Book{Title: "Titulo teste", Author: "jeffotoni", Pages: 200, Chapters: []string{"C", "C++", "Algol", "Basic"}}
        b, err := json.Marshal(book)
        if err != nil {
            log.Println(err)
        }
        io.WriteString(w, string(b))
    }
}

func BenchmarkWriteJsonMarshal(b *testing.B) {
    for n := 0; n < b.N; n++ {
        w := httptest.NewRecorder()
        var book = &Book{Title: "Titulo teste", Author: "jeffotoni", Pages: 200, Chapters: []string{"C", "C++", "Algol", "Basic"}}
        b, err := json.Marshal(book)
        if err != nil {
            log.Println(err)
            continue
        }
        w.Write(b)
    }
}

func BenchmarkNativoNewDecoderNewDecoder(b *testing.B) {
    route := http.NewServeMux()
    route.HandleFunc("/v1/api/customer", PostNativoNewDecoderNewDecoder)
    for n := 0; n < b.N; n++ {
        w := httptest.NewRecorder()
        req := httptest.NewRequest("POST", "/v1/api/customer", bytes.NewBuffer([]byte(jsonExample)))
        req.Header.Set("Content-Type", "application/json")
        //PostNativoNewDecoderNewDecoder(w, req)
        route.ServeHTTP(w, req)
        resp := w.Result()
        defer resp.Body.Close()
        _, _ = ioutil.ReadAll(resp.Body)
    }
}

func PostNativoNewDecoderNewDecoder(w http.ResponseWriter, r *http.Request) {
    var gloss = &Glossary{}
    err := json.NewDecoder(r.Body).Decode(gloss)
    if err != nil {
        log.Println(err)
        return
    }

    var book = &Book{Title: "Titulo teste", Author: "jeffotoni", Pages: 200, Chapters: []string{"C", "C++", "Algol", "Basic"}}
    err = json.NewEncoder(w).Encode(book)
    if err != nil {
        log.Println(err)
    }
}

func BenchmarkNativoNewDecoderWrite(b *testing.B) {
    route := http.NewServeMux()
    route.HandleFunc("/v1/api/customer", PostNativoMarshalWrite)
    for n := 0; n < b.N; n++ {
        w := httptest.NewRecorder()
        req := httptest.NewRequest("POST", "/v1/api/customer", bytes.NewBuffer([]byte(jsonExample)))
        req.Header.Set("Content-Type", "application/json")
        route.ServeHTTP(w, req)
        resp := w.Result()
        defer resp.Body.Close()
        var book Book
        json.NewDecoder(resp.Body).Decode(&book)
        //_, _ = ioutil.ReadAll()
    }
}

func PostNativoMarshalWrite(w http.ResponseWriter, r *http.Request) {
    var gloss = &Glossary{}
    err := json.NewDecoder(r.Body).Decode(gloss)
    if err != nil {
        log.Println(err)
        return
    }

    // ... do gloss database etc..

    var book = &Book{Title: "Titulo teste", Author: "jeffotoni", Pages: 200, Chapters: []string{"C", "C++", "Algol", "Basic"}}
    b, err := json.Marshal(book)
    if err != nil {
        log.Println(err)
        return
    }
    w.Write(b)
}

func BenchmarkFiberJson(b *testing.B) {
    app := fiber.New(&fiber.Settings{
        Prefork: true,
    })
    app.Post("/v1/api/customer", PostFiber)
    for n := 0; n < b.N; n++ {
        req := httptest.NewRequest("POST", "/v1/api/customer", bytes.NewBuffer([]byte(jsonExample)))
        req.Header.Set("Content-Type", "application/json")
        resp, _ := app.Test(req)
        _, _ = ioutil.ReadAll(resp.Body)
        defer resp.Body.Close()
        _, _ = ioutil.ReadAll(resp.Body)
    }
}

func PostFiber(c *fiber.Ctx) {
    var gloss = &Glossary{}
    err := c.BodyParser(gloss)
    if err != nil {
        log.Println(err)
        return
    }

    // ... do gloss database etc..

    var book = &Book{Title: "Titulo teste", Author: "jeffotoni", Pages: 200, Chapters: []string{"C", "C++", "Algol", "Basic"}}
    c.Status(200).JSON(book)
}
