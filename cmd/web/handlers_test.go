package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const mockEmail = "john.doe@gmail.com"

func TestMockGetPerson(t *testing.T) {
	app := newTestApplication(t)

	ts := httptest.NewServer(app.routes())
	defer ts.Close()

	r := newGetRequest(t, ts.URL+"/testImage", app.validAuthHeader)
	rs, err := ts.Client().Do(r)
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

func TestGetEncodingStringByMlResponse(t *testing.T) {
	response := []byte("[[1 2 3]]")

	encodingString := getEncodingStringByMlResponse(response)

	const correctEncodingString = "1 2 3"
	if encodingString != correctEncodingString {
		t.Errorf("want '%q', got %q", correctEncodingString, encodingString)
	}
}
