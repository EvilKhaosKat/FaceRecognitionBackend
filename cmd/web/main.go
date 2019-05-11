package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"romangaranin.dev/FaceRecognitionBackend/pkg/models/mongodb"
	"time"
)

type application struct {
	errorLog        *log.Logger
	infoLog         *log.Logger
	persons         *mongodb.PersonModel
	validAuthHeader string
}

var timeoutCtx, _ = context.WithTimeout(context.Background(), 7*time.Second)

func main() {
	dsn := flag.String("dsn", "mongodb://mongo:27017", "MongoDB data source name")
	addr := flag.String("addr", ":10080", "HTTP network address")
	validAuthHeader := flag.String("validAuthHeader", "", "Valid auth header for mock auth logic")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	infoLog.Println("Connecting to MongoDB")
	client, err := mongodb.OpenDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer client.Disconnect(timeoutCtx)

	app := &application{
		errorLog:        errorLog,
		infoLog:         infoLog,
		persons:         mongodb.NewPersonModel(client),
		validAuthHeader: *validAuthHeader,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting HTTP server on %s", *addr)
	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}
