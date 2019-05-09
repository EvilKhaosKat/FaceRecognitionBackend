package models

import "errors"

var ErrNoRecord = errors.New("models: no matching record found")

//        "first_name": "John",
//        "last_name": "Doe",
//        "email": "john.doe@gmail.com",
//        "id": "john.doe",
//        "activations": [0.09, 0.93, 0.777],
//        "confidence": 0.9

//Person
type Person struct {
	FirstName string `json:"first_name" bson:"firstName"`
	LastName  string `json:"last_name" bson:"lastName"`
	Email     string `json:"email"`
	ID        string `json:"id"`
	//TODO activations
	//TODO images
}
