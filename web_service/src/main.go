package main

import (
	"api"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"log"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	logFile := "server_debug.log"
	api.OpenLogFile(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println("Web service for testing GO")

	// lazily open db (doesn't truly open until first request)
	dbLogic, err := sql.Open("mysql", "root:pass@tcp(db:3306)/GOTEST")
	checkErr(err)
	defer dbLogic.Close()
	dbLogic.SetMaxOpenConns(10) // not unlimited number of connections

	// lazily open db (doesn't truly open until first request)
	dbStats, err := sql.Open("mysql", "root:pass@tcp(db:3306)/CDR")
	checkErr(err)
	defer dbStats.Close()
	dbStats.SetMaxOpenConns(10) // not unlimited number of connections

	log.Fatal(http.ListenAndServe(":8000", api.Handlers(dbLogic, dbStats)))
}
