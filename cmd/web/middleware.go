package main

import (
	"net/http"
	"strings"
)

//TODO implement real one
func (app *application) authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if strings.EqualFold(auth, app.validAuthHeader) {
			next.ServeHTTP(w, r)
		} else {
			app.errorLog.Printf("Not authorized call detected: Authorization:%s\nRequest:%v \n", auth, r)
		}
	})
}

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

//TODO add global id per request
