package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lalanne/GoExperiment/web_service/src/api"
	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
	"io"
	"log"
	"net/http"
)

// Init returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func tracer_init(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	logFile := "server_debug.log"
	api.OpenLogFile(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	tracer, closer := tracer_init("hello-world")
	defer closer.Close()

	span := tracer.StartSpan("request")
	span.SetTag("hello-to", "loco")

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
