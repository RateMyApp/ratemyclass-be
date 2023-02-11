package dao

import (
	"context"
	"log"

	"github.com/ratemyapp/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDao struct {
	Client *mongo.Client
}

func NewMongoDao(appConfig config.AppConfig) *MongoDao {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(appConfig.MONGO_URI))
	if err != nil {
		log.Fatal("Could not connect to mongo database: " + err.Error())
	}

	return &MongoDao{Client: client}
}
