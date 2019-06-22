package mock

import "github.com/EvilKhaosKat/FaceRecognitionBackend/pkg/models"

var mockPerson = &models.Person{
	FirstName: "First",
	LastName:  "Name",
	Email:     "email@email.com",
	ID:        "1",
	Encodings: nil,
}

type PersonsModel struct {
}

func (*PersonsModel) Update(id, firstName, lastName, email string, rawEncodings []string) (string, error) {
	panic("implement me")
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
	panic("implement me")
}

func (*PersonsModel) GetAll() ([]*models.Person, error) {
	panic("implement me")
}
