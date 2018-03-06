package main

import (
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

func main() {
	logFile := "server_debug.log"
	handlers.openLogFile(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Web service for testing GO")

	router := mux.NewRouter()
	router.HandleFunc("/", handlers.genericHandler).Methods("GET")
	router.HandleFunc("/purchase", handlers.purchaseHandler).Methods("GET")
	router.HandleFunc("/sale", handlers.saleHandler).Methods("GET")

	/*starts an http server*/
	log.Fatal(http.ListenAndServe(":8000", router))
}
