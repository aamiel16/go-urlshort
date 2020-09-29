package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aamiel16/go-urlshort"
)

func main() {
	mux := defaultMux()

	handler, err := urlshort.RedirectHandler("./mapping.json", mux)
	// handler, err := urlshort.YAMLHandler("./mapping.yml", mux)
	if err != nil {
		panic(err)
	}

	const PORT = 8080
	fmt.Println("Server started. Listening on port", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), handler))
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
