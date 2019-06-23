package models

import "errors"

var ErrNoRecord = errors.New("models: no matching record found")
var ErrDbProblem = errors.New("models: problem with db")

//Person
type Person struct {
	FirstName string   `json:"first_name" bson:"firstName"`
	LastName  string   `json:"last_name" bson:"lastName"`
	Email     string   `json:"email" bson:"email"`
	ID        string   `json:"id" bson:"id"`
	Encodings []string `json:"encodings" bson:"encodings"`
	//TODO images
	//TODO confidence level?
}

//PersonModel defines model/DAO methods for Person
type PersonModel interface {
	Update(id, firstName, lastName, email string, encodings []string) (string, error)
	Get(id string) (*Person, error)
	Remove(id string) (int64, error)
	GetAll() ([]*Person, error)
}
