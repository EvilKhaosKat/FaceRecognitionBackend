package services

import (
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

	for _, person := range persons {
		for _, rawEncoding := range person.RawEncodings {
			personEncoding, err := NewEncoding(rawEncoding)
			if err != nil {
				return nil, err
			}

			samePerson, err := encoding.IsSame(personEncoding)
			if err != nil {
				return nil, err
			}
			if samePerson {
				return person, nil
			}
		}
	}

	return nil, nil
}
