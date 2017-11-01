package main

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

func openLogFile(logfile string) {
	if logfile != "" {
		lf, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

		if err != nil {
			log.Fatal("OpenLogfile: os.OpenFile:", err)
		}

		log.SetOutput(lf)
	}
}

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

	logFile := "development.log"
	openLogFile(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Web service for testing GO")

	router := mux.NewRouter()
	router.HandleFunc("/", generic_handler).Methods("GET")
	router.HandleFunc("/purchase", purchase_handler).Methods("GET")
	router.HandleFunc("/sale", sale_handler).Methods("GET")

	/*starts an http server*/
	log.Fatal(http.ListenAndServe(":8000", router))
}
