package repositories

import "go.uber.org/fx"

var Module = fx.Module("repositories", fx.Provide(
	fx.Annotate(
		NewUserRepository, fx.As(new(UserRepository)),
	),
	fx.Annotate(
		NewProfessorRepository, fx.As(new(ProfessorRepository)),
	),
	fx.Annotate(
		NewCourseRatingRepository, fx.As(new(CourseRatingRepository)),
	),
	fx.Annotate(NewCoursRepository, fx.As(new(CourseRepository)),
)))
