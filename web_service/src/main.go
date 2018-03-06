package main

import (
	"api"
	"log"
	"net/http"
)

func main() {
	logFile := "server_debug.log"
	api.OpenLogFile(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println("Web service for testing GO")

	log.Fatal(http.ListenAndServe(":8000", api.Handlers()))
}
