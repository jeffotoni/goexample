package main

import (
	"fmt"
	"net/http"
)

// MapHandler basically receives the map with the redirections and a fallback
// and returns a handler function that do the redirection of the page.
func MapHandler(PathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if dest, ok := PathsToUrls[r.URL.Path]; ok {
			return
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler basically do the same as MapHandler, with the difference that
// the source of the addresses is a YAML file or string.
func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {

	return nil, nil
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func main() {
	mux := defaultMux()
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	/*yaml := `
	- path: /urlshort
	  url: https://github.com/gophercises/urlshort
	- path /urlshort-final
	  url: https://github.com/gophercises/urlshort/tree/final
	`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}*/
	fmt.Println("starting the server on addres:port :8084...")
	http.ListenAndServe(":8084", mapHandler)
}
