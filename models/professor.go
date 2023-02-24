package models

type Professor struct {
	Name         string
	Email        string
	Department   string
	Courses      []*Course `gorm:"many2many:professor_courses"`
	CourseRating []CourseRating
	ProgramID    uint
}
