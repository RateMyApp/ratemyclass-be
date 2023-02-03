package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type User struct{
	ID primitive.ObjectID `bson:"_id"`
	FirstName string
	LastName string
	Email string 
	Password string

}