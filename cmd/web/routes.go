package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/person", app.getPerson)
	mux.HandleFunc("/person/create", app.createPerson)
	mux.HandleFunc("/testImage", app.mockGetPerson)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//TODO use alice middleware?
	return app.logRequest(app.authorization(mux, app.validAuthHeader))
}
