package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"romangaranin.dev/FaceRecognitionBackend/pkg/models"
	"time"
)

var ctx, _ = context.WithTimeout(context.Background(), 7*time.Second)

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
	DB *mongo.Client
}

// This will insert a new person into the database.
func (m *PersonModel) Insert(id, firstName, lastName, email string) (int, error) {
	return 0, nil
}

// This will return a specific person based on its id.
func (m *PersonModel) Get(id int) (*models.Person, error) {
	return nil, nil
}

// This will return all the created persons.
func (m *PersonModel) GetAll() ([]*models.Person, error) {
	return nil, nil
}
