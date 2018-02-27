package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
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

func getSoapInfo() {
	log.Printf("[getSoapInfo] \n")

}

func validateOperation(w http.ResponseWriter, c chan int) {
	log.Printf("[validateOperation]\n")
	// lazily open db (doesn't truly open until first request)
	db, err := sql.Open("mysql", "root:pass@tcp(db:3306)/GOTEST")
	checkErr(err)
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*1)
	defer cancel()

	rows, err := db.QueryContext(
		ctx,
		"select count(1) from OperationsAllowed where op=\"purchase\";",
	)
	checkErr(err)
	defer rows.Close()

	var count int
	rows.Next()
	err = rows.Scan(&count)
	checkErr(err)

	log.Printf("[validateOperation] count[%d]\n", count)

	if count == 1 {
		c <- 0
	} else {
		c <- 1
	}
}

func insertCdr(w http.ResponseWriter, c chan int) {
	log.Printf("[insertCdr]\n")
	// lazily open db (doesn't truly open until first request)
	db, err := sql.Open("mysql", "root:pass@tcp(db:3306)/CDR")
	checkErr(err)
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*1)
	defer cancel()

	res, err := db.ExecContext(
		ctx,
		"insert into cdr (id, error, host, op) values (\"1\", 0, 0, \"purchase\");",
	)
	checkErr(err)

	log.Printf("[insertCdr] res[%d]\n", res)

	c <- 0
}

func purchaseHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[Purchase Handler] [%s] [%s] [%s]\n", r.RemoteAddr, r.Method, r.URL)
	c := make(chan int)
	c1 := make(chan int)

	go validateOperation(w, c)
	x := <-c
	log.Printf("[Purchase Handler] return from validateOperation success? [%d]\n", x)

	go insertCdr(w, c1)
	x1 := <-c1
	log.Printf("[Purchase Handler] return from insertCdr success? [%d]\n", x1)

	io.WriteString(w, "purchase operation allowed by DB!")
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
