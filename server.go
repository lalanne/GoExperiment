package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func generic_handler(w http.ResponseWriter, r *http.Request) {
	log.Println("generic handler")

	io.WriteString(w, "hello world!")
}

func purchase_handler(w http.ResponseWriter, r *http.Request) {
	log.Println("purchase handler")

	io.WriteString(w, "hello world!")
}

func sale_handler(w http.ResponseWriter, r *http.Request) {
	log.Println("sale handler")

	io.WriteString(w, "hello world!")
}

func main() {
	fmt.Println("vim-go")

	router := mux.NewRouter()
	router.HandleFunc("/", generic_handler).Methods("GET")
	router.HandleFunc("/purchase", purchase_handler).Methods("GET")
	router.HandleFunc("/sale", sale_handler).Methods("GET")

	/*starts an http server*/
	log.Fatal(http.ListenAndServe(":8000", router))
}
