package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"romangaranin.dev/FaceRecognitionBackend/pkg/models"
	"romangaranin.dev/FaceRecognitionBackend/pkg/services"
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

	_, err = app.persons.Update(p.ID, p.FirstName, p.LastName, p.Email, p.RawEncodings)
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

func (app *application) getImageRawEncoding(img io.Reader) (services.Encoding, error) {
	response, err := httpClient.Post(app.mlEndpoint, "image/jpeg", img)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	encoding, err := services.NewEncoding(string(responseBody))
	if err != nil {
		return nil, err
	}

	return encoding, nil
}

func (app *application) checkPerson(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("Check person called")
	img, err := ioutil.ReadAll(r.Body)
	if err != nil {
		app.serverError(w, err)
		return
	}

	encoding, err := app.getImageRawEncoding(bytes.NewReader(img))
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.infoLog.Println("ML response obtained")

	foundPerson, err := app.encodingChecker.FindSamePerson(encoding)
	if err != nil {
		app.serverError(w, err)
	}

	if foundPerson != nil {
		err = json.NewEncoder(w).Encode(foundPerson)
	} else {
		_, err = fmt.Fprint(w, "{}")
	}
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) deletePerson(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.serverError(w, err)
		return
	}

	personId := r.FormValue("id")
	person, err := app.persons.Get(personId)
	if err != nil {
		app.errorLog.Println(err)
		app.notFound(w)
		return
	}

	if person == nil {
		app.notFound(w)
		return
	}

	removedCount, err := app.persons.Remove(person.ID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	_, err = fmt.Fprintf(w, "Removed:%d", removedCount)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) addImageToPerson(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(5 * 1024 * 1025)
	if err != nil {
		log.Println(err)
		return
	}

	personId := r.FormValue("id")
	person, err := app.persons.Get(personId)
	if err != nil {
		app.errorLog.Println(err)
		app.notFound(w)
		return
	}

	if person == nil {
		app.notFound(w)
		return
	}

	img, _, err := r.FormFile("image")
	if err != nil {
		app.serverError(w, err)
		return
	}
	defer img.Close()

	imgBuf := bytes.NewBuffer(nil)
	if _, err := io.Copy(imgBuf, img); err != nil {
		app.serverError(w, err)
		return
	}

	encoding, err := app.getImageRawEncoding(imgBuf)
	if err != nil {
		app.serverError(w, err)
		return
	}

	//TODO get rid of that 'raw encoding' form
	rawEncodingForm := "[" + fmt.Sprint(encoding) + "]"

	person.RawEncodings = append(person.RawEncodings, rawEncodingForm)
	_, err = app.persons.Update(person.ID, person.FirstName, person.LastName, person.Email, person.RawEncodings)
	if err != nil {
		app.serverError(w, err)
		return
	}
}
