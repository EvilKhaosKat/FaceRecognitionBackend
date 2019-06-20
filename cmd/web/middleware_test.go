package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

const authHeader = "AUTH_HEADER"

var next = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
})

func TestAuthValidHeader(t *testing.T) {
	//given
	rr := httptest.NewRecorder()
	r := newRequest(t, authHeader)

	app := newMockApp(authHeader)

	//when
	app.authorization(next).ServeHTTP(rr, r)

	//then
	rs := rr.Result()
	checkStatusCodeOk(t, rs)

	defer rs.Body.Close()
	body := readBody(t, rs)
	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}
}

func TestAuthNotValidHeader(t *testing.T) {
	//given
	rr := httptest.NewRecorder()
	r := newRequest(t, "wrong auth header")

	app := newMockApp(authHeader)

	//when
	app.authorization(next).ServeHTTP(rr, r)

	//then
	rs := rr.Result()
	checkStatusCodeOk(t, rs)

	defer rs.Body.Close()
	body := readBody(t, rs)
	if string(body) == "OK" {
		t.Errorf("want body to not equal %q", "OK")
	}
}

func readBody(t *testing.T, rs *http.Response) []byte {
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	return body
}

func checkStatusCodeOk(t *testing.T, rs *http.Response) {
	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	}
}

func newRequest(t *testing.T, authHeader string) *http.Request {
	r, err := http.NewRequest("GET", "/", nil)
	r.Header.Set("Authorization", authHeader)
	if err != nil {
		t.Fatal(err)
	}
	return r
}

func newMockApp(authHeader string) *application {
	app := &application{
		errorLog:        log.New(ioutil.Discard, "", 0),
		infoLog:         log.New(ioutil.Discard, "", 0),
		validAuthHeader: authHeader,
	}
	return app
}
