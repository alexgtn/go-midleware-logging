package main

import (
	"fmt"
	"github.com/alexgtn/go-midleware-logging/log"
	"github.com/alexgtn/go-midleware-logging/middleware"
	"github.com/gorilla/mux"
	"net"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	conn, err := net.Dial("tcp", "localhost:5555")
	if err != nil {
		fmt.Errorf("error: %s", err.Error())
		return
	}

	logger := log.NewGraylogLogger(conn)
	//logger := logrus.New()
	loggingMiddleware := middleware.NewLoggingMiddleware(logger)
	r.HandleFunc("/lemon", lemonHandler).Methods(http.MethodGet)
	r.HandleFunc("/potato", potatoHandler).Methods(http.MethodPost)
	r.Use(loggingMiddleware.Logging)

	http.ListenAndServe(":8080", r)
}

func lemonHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Lemon"))
}

func potatoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Potato"))
}

