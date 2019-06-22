package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const mockEmail = "john.doe@gmail.com"

func TestMockGetPerson(t *testing.T) {
	app := &application{
		errorLog: log.New(ioutil.Discard, "", 0),
		infoLog:  log.New(ioutil.Discard, "", 0),
	}

	ts := httptest.NewServer(app.routes())
	defer ts.Close()

	rs, err := ts.Client().Get(ts.URL + "/testImage")
	if err != nil {
		t.Fatal(err)
	}

	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	}

	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(body), mockEmail) {
		t.Errorf("want body contains json with email %q", mockEmail)
	}
}
