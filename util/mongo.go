package util

import (
	"comunty/ms-auth/conf"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
)

type Mongo struct {
	DB         string
	Collection string
}

func init() {
	ctx := context.TODO()
	config := conf.GlobalConf.Credentials.Persist.MongoDB
	connURI := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", config.User, config.Secret, config.Host)
	clientOptions := options.Client().ApplyURI(connURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (m *Mongo) SetCollection() *mongo.Collection {
	return client.Database(m.DB).Collection(m.Collection)
}

func (m *Mongo) InsertOne(ctx context.Context, doc interface{}) (*mongo.InsertOneResult, error) {
	return client.Database(m.DB).Collection(m.Collection).InsertOne(ctx, doc)
}

func (m *Mongo) Find(filters interface{}, opts *options.FindOptions, ctx context.Context) (*mongo.Cursor, error) {
	return client.Database(m.DB).Collection(m.Collection).Find(ctx, bson.D{})
}
