package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetbox"))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", home)

	log.Print("Starting server on port :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
