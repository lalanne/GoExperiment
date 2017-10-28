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

	/*	decoder := json.NewDecoder(r.Body)
		var req Request
		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}*/

	io.WriteString(w, "hello world!")
}

func main() {
	fmt.Println("vim-go")

	router := mux.NewRouter()
	router.HandleFunc("/", handler).Methods("GET")
	http.ListenAndServe(":8000", router)
}
