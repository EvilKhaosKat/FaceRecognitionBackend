package main

import (
	"github.com/EvilKhaosKat/FaceRecognitionBackend/pkg/models/mock"
	"github.com/EvilKhaosKat/FaceRecognitionBackend/pkg/services"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

const mockMlEndpoint = "localhost:4242/ml"
const mockValidAuthHeader = "test"

func newTestApplication(t *testing.T) *application {
	personsModel := &mock.PersonsModel{}
	return &application{
		errorLog:           log.New(ioutil.Discard, "", 0),
		infoLog:            log.New(ioutil.Discard, "", 0),
		persons:            personsModel,
		encodingComparator: services.NewEncodingComparator(personsModel),
		validAuthHeader:    mockValidAuthHeader,
		mlEndpoint:         mockMlEndpoint,
	}
}

func newGetRequest(t *testing.T, url, authHeader string) *http.Request {
	r, err := http.NewRequest("GET", url, nil)
	r.Header.Set("Authorization", authHeader)
	if err != nil {
		t.Fatal(err)
	}
	return r
}
