package logger_test

import (
	"net/http"
	"os"

	"github.com/go-http-utils/logger"
)

func ExampleLogger_DefaultHandler() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":8080", logger.DefaultHandler(mux))
}

func ExampleLogger_Handler() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":8080", logger.Handler(os.Stdout, logger.DevLoggerType, mux))
}