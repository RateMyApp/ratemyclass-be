package models

import "database/sql/driver"

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

// Enums
type status string

const (
	PENDING  status = "PENDING"
	APPROVED status = "APPROVED"
	DECLINED status = "DECLINED"
)

func (self *status) Scan(value interface{}) error {
	*self = status(value.([]byte))
	return nil
}

func (self status) Value() (driver.Value, error) {
	return string(self), nil
}
