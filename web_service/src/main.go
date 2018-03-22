package main

import (
	"api"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	//"github.com/uber/jaeger-lib/metrics"
	"log"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	jLogger := jaegerlog.StdLogger
	//jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	closer, err := cfg.InitGlobalTracer(
		"orquestador",
		jaegercfg.Logger(jLogger),
	//	jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	defer closer.Close()

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
