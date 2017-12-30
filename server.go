package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func openLogFile(logfile string) {
	if logfile != "" {
		lf, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)
		checkErr(err)

		log.SetOutput(lf)
	}
}

func genericHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[Generic Handler] [%s] [%s] [%s]\n", r.RemoteAddr, r.Method, r.URL)

	io.WriteString(w, "default operation!")
}

func purchaseHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[Purchase Handler] [%s] [%s] [%s]\n", r.RemoteAddr, r.Method, r.URL)

	db, err := sql.Open("mysql", "root:pass@tcp(0.0.0.0:3306)/GOTEST")
	checkErr(err)
	defer db.Close()

	/*Prepare statement for reading data*/
	rows, err := db.Query("select * from OperationsAllowed")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var id string
		var error int
		var host int
		var op string
		err := rows.Scan(&id, &error, &host, &op)
		checkErr(err)

		if op == "purchase" {
			io.WriteString(w, "purchase operation allowed by DB!")
			break
		} else {
			io.WriteString(w, "purchase operation NOT allowed by DB!")
		}
	}

}

func saleHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[Sale Handler] [%s] [%s] [%s]\n", r.RemoteAddr, r.Method, r.URL)

	s, err := regexp.Compile(`\?(.*)`)
	checkErr(err)

	res := s.FindAllString(r.URL.String(), -1)
	log.Printf("<%v>\n", res)

	io.WriteString(w, "sale operation!")
}

func main() {
	logFile := "server_debug.log"
	openLogFile(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Web service for testing GO")

	router := mux.NewRouter()
	router.HandleFunc("/", genericHandler).Methods("GET")
	router.HandleFunc("/purchase", purchaseHandler).Methods("GET")
	router.HandleFunc("/sale", saleHandler).Methods("GET")

	/*starts an http server*/
	log.Fatal(http.ListenAndServe(":8000", router))
}
