package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"romangaranin.dev/FaceRecognitionBackend/pkg/models/mongodb"
	"romangaranin.dev/FaceRecognitionBackend/pkg/services"
	"time"
)

type application struct {
	errorLog        *log.Logger
	infoLog         *log.Logger
	persons         *mongodb.PersonModel
	encodingChecker *services.EncodingChecker
	validAuthHeader string //TODO move to separate config struct
	mlEndpoint      string
}

var timeoutCtx, _ = context.WithTimeout(context.Background(), 7*time.Second)

func main() {
	dsn := flag.String("dsn", "mongodb://localhost:27017", "MongoDB data source name")
	mlEndpoint := flag.String("mlEndpoint", "http://localhost:8000", "Machine learning model endpoint")
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

	personModel := mongodb.NewPersonModel(client)
	encodingChecker := services.NewEncodingChecker(personModel)

	app := &application{
		errorLog:        errorLog,
		infoLog:         infoLog,
		persons:         personModel,
		encodingChecker: encodingChecker,
		validAuthHeader: *validAuthHeader,
		mlEndpoint:      *mlEndpoint,
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
