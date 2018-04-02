package api

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
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

func OpenLogFile(logfile string) {
	if logfile != "" {
		lf, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)
		checkErr(err)

		log.SetOutput(lf)
	}
}

func getSoapInfo() {
	log.Printf("[getSoapInfo] \n")

}

func validateOperation(w http.ResponseWriter,
	c chan int,
	db *sql.DB) {
	log.Printf("[validateOperation]\n")

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

func insertCdr(w http.ResponseWriter,
	c chan int,
	db *sql.DB) {
	log.Printf("[insertCdr]\n")

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

func getHTTPResponse(w http.ResponseWriter, c chan int) {
	log.Printf("[getHTTPResponse]\n")

	var httpClient = &http.Client{
		Timeout: time.Second * 1,
	}
	resp, err := httpClient.Get("http://http:8080/")
	checkErr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)
	log.Printf("[getHTTPResponse][%s]\n", bodyString)

	c <- 0
}

func purchaseHandler(w http.ResponseWriter,
	r *http.Request,
	dbLogic *sql.DB,
	dbStats *sql.DB) {
	log.Printf("[Purchase Handler] [%s] [%s] [%s]\n", r.RemoteAddr, r.Method, r.URL)

	c := make(chan int)
	c0 := make(chan int)
	c1 := make(chan int)

	go validateOperation(w,
		c,
		dbLogic)
	x := <-c
	log.Printf("[Purchase Handler] return from validateOperation success? [%d]\n", x)

	go getHTTPResponse(w, c0)
	x0 := <-c0
	log.Printf("[Purchase Handler] return from getHTTPResponse success? [%d]\n", x0)

	go insertCdr(w,
		c1,
		dbStats)
	x1 := <-c1
	log.Printf("[Purchase Handler] return from insertCdr success? [%d]\n", x1)

	io.WriteString(w, "Web service answer, everything OK with happy path!!!")
}

func saleHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[Sale Handler] [%s] [%s] [%s]\n", r.RemoteAddr, r.Method, r.URL)

	s, err := regexp.Compile(`\?(.*)`)
	checkErr(err)

	res := s.FindAllString(r.URL.String(), -1)
	log.Printf("<%v>\n", res)

	io.WriteString(w, "sale operation!")
}

func genericHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[Generic Handler] [%s] [%s] [%s]\n", r.RemoteAddr, r.Method, r.URL)

	io.WriteString(w, "default operation!")
}

func Handlers(dbLogic *sql.DB, dbStats *sql.DB) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", genericHandler).Methods("GET")

	r.HandleFunc("/purchase", func(w http.ResponseWriter, r *http.Request) {
		purchaseHandler(w,
			r,
			dbLogic,
			dbStats)
	}).Methods("GET")

	r.HandleFunc("/sale", saleHandler).Methods("GET")

	return r
}
