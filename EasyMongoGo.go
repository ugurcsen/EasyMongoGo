package EasyMongoGo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EasyMongo struct {
	Client      *mongo.Client
	Database    *mongo.Database
	ContextFunc func() context.Context
}

type Collection struct {
	*mongo.Collection
	ContextFunc func() context.Context
}

//Connect MongoDB Database And Initialize EasyMongo Object
func NewEasyMongo(uri string) (*EasyMongo, error) {
	v := new(EasyMongo)
	v.Client = nil
	v.Database = nil
	v.ContextFunc = context.TODO
	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(v.ContextFunc(), clientOptions)

	if err != nil {
		return v, err
	}

	v.Client = client

	return v, nil
}

//Pinging MongoDB
func (v *EasyMongo) Ping() error {
	// Check the connection
	err := v.Client.Ping(context.TODO(), nil)

	if err != nil {
		return err
	}

	return nil
}

//Select Database
func (v *EasyMongo) UseDatabase(dbName string) {
	v.Database = v.Client.Database(dbName)
}

//Select Collection
func (v *EasyMongo) SelectCollection(collectionName string) *Collection {
	if v.Database == nil {
		return nil
	}
	return &Collection{v.Database.Collection(collectionName), v.ContextFunc}
}