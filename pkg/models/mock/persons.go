package mock

import "github.com/EvilKhaosKat/FaceRecognitionBackend/pkg/models"

var mockPerson = &models.Person{
	FirstName: "First",
	LastName:  "Name",
	Email:     "email@email.com",
	ID:        "1",
	Encodings: []string{"1 2 3"},
}

type PersonsModel struct {
}

func (*PersonsModel) Update(id, firstName, lastName, email string, encodings []string) (string, error) {
	switch id {
	case "1":
		mockPerson.ID = id
		mockPerson.FirstName = firstName
		mockPerson.LastName = lastName
		mockPerson.Email = email
		mockPerson.Encodings = encodings

		return id, nil
	default:
		return "", models.ErrDbProblem
	}
}

func (*PersonsModel) Get(id string) (*models.Person, error) {
	switch id {
	case "1":
		return mockPerson, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (*PersonsModel) Remove(id string) (int64, error) {
	return 0, nil
}

func (*PersonsModel) GetAll() ([]*models.Person, error) {
	return []*models.Person{mockPerson}, nil
}
