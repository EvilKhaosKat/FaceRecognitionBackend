package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/person/add", app.addPerson)
	mux.HandleFunc("/person/get", app.getPerson)
	mux.HandleFunc("/person/all", app.getPersons)
	mux.HandleFunc("/person/check", app.checkPerson)
	mux.HandleFunc("/testImage", app.mockGetPerson)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//TODO use alice middleware?
	return app.logRequest(app.authorization(mux, app.validAuthHeader))
}
