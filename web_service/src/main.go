package main

import (
	"api"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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

	// lazily open db (doesn't truly open until first request)
	dbStats, err := sql.Open("mysql", "root:pass@tcp(db:3306)/CDR")
	checkErr(err)
	defer dbStats.Close()

	log.Fatal(http.ListenAndServe(":8000", api.Handlers(dbLogic, dbStats)))
}
