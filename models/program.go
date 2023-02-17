package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Program struct {
	ID primitive.ObjectID

	School string 
	CourseCode string "bson:code"
	CourseName string "bson:name"

}
