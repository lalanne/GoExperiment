package http_server

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
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

func main() {
	logFile := "http_debug.log"
	openLogFile(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Http server.....")

	router := mux.NewRouter()
	router.HandleFunc("/", genericHandler).Methods("GET")

	/*starts an http server*/
	log.Fatal(http.ListenAndServe(":8080", router))
}
