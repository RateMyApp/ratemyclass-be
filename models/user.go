package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type School struct{
	ID primitive.ObjectID `bson:"id"`
	Name string `bson:"firstname"`
	Location string`bson:"location"`
	StudentHeadcount int `bson:"studentheadcount"`
	StaffHeadcount int `bson:"staffheadcount"`
}