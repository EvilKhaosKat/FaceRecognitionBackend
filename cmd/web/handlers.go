package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"romangaranin.dev/FaceRecognitionBackend/pkg/models"
)

//Mock
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
		app.errorLog.Print(err)
	}
}

func (app *application) getPerson(w http.ResponseWriter, r *http.Request) {

}

func (app *application) createPerson(w http.ResponseWriter, r *http.Request) {
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

	app.infoLog.Println("POST person:", p)
}
