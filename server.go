package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("------request arrived!-----")

	io.WriteString(w, "hello world!")
}

func main() {
	fmt.Println("vim-go")

	router := mux.NewRouter()
	router.HandleFunc("/", handler).Methods("GET")

	/*starts an http server*/
	http.ListenAndServe(":8000", router)
}
