package main

import (
	"flag"
	"fmt"
	"sample_microservice/server"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"

	"github.com/go-kit/log/level"

	"net/http"
	"os"
)

func main() {
	var httpAddr = flag.String("http", ":8080", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	flag.Parse()
	r := gin.Default()
	var srv server.Service
	{
		srv = server.NewService(logger)
	}

	errs := make(chan error)

	endpoints := server.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := server.NewHTTPServer(r, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
