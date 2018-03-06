package main

import (
	"api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	logFile := "server_debug.log"
	api.OpenLogFile(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Web service for testing GO")

	router := mux.NewRouter()
	router.HandleFunc("/", api.GenericHandler).Methods("GET")
	router.HandleFunc("/purchase", api.PurchaseHandler).Methods("GET")
	router.HandleFunc("/sale", api.SaleHandler).Methods("GET")

	/*starts an http server*/
	log.Fatal(http.ListenAndServe(":8000", router))
}
