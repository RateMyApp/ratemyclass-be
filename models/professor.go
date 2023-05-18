package models

import "gorm.io/gorm"

type Professor struct {
	gorm.Model
	Firstname        string
	Lastname         string
	DirectoryListing string
	Email            string
	Department       string
	Courses          []*Course `gorm:"many2many:professor_courses"`
	CourseRating     []CourseRating
	Status           status `sql:"type:enum('PENDING','APPROVED','DECLINED')" gorm:"column:status;default:'PENDING'"`
}
