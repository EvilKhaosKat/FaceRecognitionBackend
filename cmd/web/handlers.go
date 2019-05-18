package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"romangaranin.dev/FaceRecognitionBackend/pkg/models"
	"time"
)

var httpClient = &http.Client{
	Timeout: time.Second * 10,
}

func (app *application) mockGetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := fmt.Fprint(w,
		`{
        "first_name": "John",
        "last_name": "Doe",
        "email": "john.doe@gmail.com",
        "id": "john.doe",
        "activations": [0.09, 0.93, 0.777],
        "confidence": 0.9
    }`)

	if err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) addPerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	var p models.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	app.infoLog.Printf("POST person:%+v \n", p)

	_, err = app.persons.Update(p.ID, p.FirstName, p.LastName, p.Email, p.RawActivations)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Println("Person added in db")
}

func (app *application) getPersons(w http.ResponseWriter, r *http.Request) {
	persons, err := app.persons.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(persons)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) getPerson(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.serverError(w, err)
		return
	}

	//TODO impl
	//rawId := r.FormValue("id")
	//person, err := app.persons.Get(rawId)
	//if err != nil {
	//
	//}
}

func (app *application) checkPerson(w http.ResponseWriter, r *http.Request) {
	img, err := ioutil.ReadAll(r.Body)
	if err != nil {
		app.serverError(w, err)
		return
	}

	//TODO hardcoded value - at least add parametrization
	response, err := httpClient.Post("http://localhost:8000", "image/jpeg", bytes.NewReader(img))
	if err != nil {
		app.serverError(w, err)
		return
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		app.serverError(w, err)
		return
	}

	//TODO use encoding to check whether there is similar person
	_, err = w.Write(responseBody)
	if err != nil {
		app.serverError(w, err)
		return
	}
}
