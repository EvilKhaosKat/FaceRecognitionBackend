package services

import (
	"math"
	"romangaranin.dev/FaceRecognitionBackend/pkg/models"
	"romangaranin.dev/FaceRecognitionBackend/pkg/models/mongodb"
)

type EncodingChecker struct {
	persons *mongodb.PersonModel
}

func NewEncodingChecker(persons *mongodb.PersonModel) *EncodingChecker {
	return &EncodingChecker{persons: persons}
}

func (e *EncodingChecker) FindSamePerson(encoding Encoding) (*models.Person, error) {
	persons, err := e.persons.GetAll()
	if err != nil {
		return nil, err
	}

	var closestPerson *models.Person
	var closestDist = math.MaxFloat64

	for _, person := range persons {
		for _, rawEncoding := range person.RawEncodings {
			personEncoding, err := NewEncoding(rawEncoding)
			if err != nil {
				return nil, err
			}

			samePerson, dist, err := encoding.IsSame(personEncoding)
			if err != nil {
				return nil, err
			}
			if samePerson && dist < closestDist {
				closestPerson = person
			}
		}
	}

	return closestPerson, nil
}
