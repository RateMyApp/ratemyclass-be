package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type School struct{
	ID primitive.ObjectID `bson:"id"`
	Name string `json:"firstname"`
	Location string`json:"location"`
	StudentHeadcount int `json:"studentheadcount"`
	StaffHeadcount int `json:"staffheadcount"`
}