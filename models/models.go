package models

func GetModels() *[]interface{} {
	return &[]interface{}{
		&School{},
		&Program{},
		&User{},
		&Course{},
		&CourseRating{},
		&Professor{},
	}
}
