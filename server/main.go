package main

import (
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello there wat!\n"))
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler)

	log.Fatal(http.ListenAndServe(":8000", r))
}
