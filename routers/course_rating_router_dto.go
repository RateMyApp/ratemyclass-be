package routers

import (
	validation "github.com/go-ozzo/ozzo-validation"
	//"github.com/go-ozzo/ozzo-validation/is"
)

type CreateCourseRatingReq struct {
	ProfessorID      uint    `json:"professorid"`
	ExperienceRating float32 `json:"experiencerating"`
	DifficultyRating float32 `json:"difficultyrating"`
	Review           string  `json:"review"`
	CourseID         uint    `json:"courseid"`
	UserID           uint    `json:"userid"`
	IsAnonymous      bool    `json:"isanonymous"`
}

// implement Validatable interface by implementing Validate
func (ccrr CreateCourseRatingReq) Validate() error {
	return validation.ValidateStruct(&ccrr,
		validation.Field(&ccrr.ProfessorID, validation.Required),
		validation.Field(&ccrr.ExperienceRating, validation.Required),
		validation.Field(&ccrr.DifficultyRating, validation.Required),
		validation.Field(&ccrr.Review, validation.Required),
		validation.Field(&ccrr.CourseID, validation.Required),
		validation.Field(&ccrr.UserID, validation.Required),
		validation.Field(&ccrr.IsAnonymous, validation.Required),
	)
}
