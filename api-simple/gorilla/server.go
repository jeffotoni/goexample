package main

import (
    _ "encoding/json"
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "github.com/gorilla/mux"
    "log"
    "math/rand"
    "net/http"
    "time"
)

type DadosProcesso struct {
    NumeroProcesso string `json:"numero-processo,omitempty"`
    Secao          string `json:"secao,omitempty"`
    Nome           string `json:"nome,omitempty"`
}

var dados []DadosProcesso

func configurarRotas() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/pegar-processo/{numeroProcesso}/{secao}/{nome}", PegarDadosProcesso).Methods("GET")

    return router
}

func subirServidor() {
    r := configurarRotas()

    porta := ":1337"
    fmt.Println("O servidor está rodando na porta:" + porta)
    log.Fatal(http.ListenAndServe(porta, r))
}

func PegarDadosProcesso(w http.ResponseWriter, r *http.Request) {
    parametros := mux.Vars(r)
    numeroProcesso := parametros["numeroProcesso"]
    secao := parametros["secao"]
    nome := parametros["nome"]

    //pessoa = append(pessoa, Pessoa{NumeroProcesso: numeroProcesso, Secao: secao, Nome: nome})
    //fmt.Println("número do processo:" + numeroProcesso + ", seção:" + secao + ", nome:" + nome)
    //json.NewEncoder(w).Encode(pessoa)

    UrlBase := "https://processual.trf1.jus.br/consultaProcessual/processo.php?proc=" + numeroProcesso + "&secao=" + secao + "&nome=" + nome + "&mostrarBaixados=N"
    fmt.Println("Url:\n", UrlBase)
    res := tratarUrl(UrlBase)
    html, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("print here...")
    fmt.Println(html.Html())
}

func tratarUrl(UrlBase string) *http.Response {
    time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
    //verificações básicas de erro a seguir
    res, err := http.Get(UrlBase)
    if err != nil {
        panic(err)
    }

    if res.StatusCode != 200 {
        panic(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status))
    }

    return res
}

func main() {
    subirServidor()
}
