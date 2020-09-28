package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aamiel16/go-urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// fallback
	//	yaml := `
	//- path: /urlshort
	//  url: https://github.com/gophercises/urlshort
	//- path: /urlshort-final
	//  url: https://github.com/gophercises/urlshort/tree/solution
	//`
	//yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	//if err != nil {
	//	panic(err)
	//}
	fmt.Println("Starting the server on :8080")
	log.Fatal(http.ListenAndServe(":8081", mapHandler))
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
