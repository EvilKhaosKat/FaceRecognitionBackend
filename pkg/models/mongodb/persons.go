package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"romangaranin.dev/FaceRecognitionBackend/pkg/models"
	"unicode/utf8"
)

const (
	databaseName      = "faceRecognition"
	collectionPersons = "persons"
)

var (
	ctx = context.Background()
)

func OpenDB(dsn string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}

type PersonModel struct {
	client *mongo.Client
}

func NewPersonModel(client *mongo.Client) *PersonModel {
	return &PersonModel{client}
}

func (m *PersonModel) getPersonsCollection() *mongo.Collection {
	return m.client.Database(databaseName).Collection(collectionPersons)
}

// This will insert a new person into the database or updates existing.
func (m *PersonModel) Update(id, firstName, lastName, email string, rawEncodings []string) (string, error) {
	persons := m.getPersonsCollection()

	upsert := true
	result, err := persons.UpdateOne(ctx,
		bson.M{"id": id},
		bson.M{
			"$set": bson.M{
				"id":        id,
				"firstName": firstName,
				"lastName":  lastName,
				"email":     email,
				"encodings": rawEncodings},
		},
		&options.UpdateOptions{
			Upsert: &upsert,
		},
	)
	if err != nil {
		return "", err
	}

	return fmt.Sprint(result.UpsertedID), nil
}

// This will return a specific person based on its id.
func (m *PersonModel) Get(id string) (*models.Person, error) {
	if utf8.RuneCountInString(id) == 0 {
		return nil, nil
	}

	persons := m.getPersonsCollection()

	result := persons.FindOne(ctx, bson.M{"id": id})

	var person *models.Person
	err := result.Decode(&person)
	if err != nil {
		return nil, err
	}

	return person, nil
}

// This will return all the created persons.
func (m *PersonModel) GetAll() ([]*models.Person, error) {
	var result []*models.Person

	persons := m.getPersonsCollection()
	cur, err := persons.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var person models.Person
		err := cur.Decode(&person)
		if err != nil {
			return nil, err
		}

		result = append(result, &person)
	}
	return result, nil
}
