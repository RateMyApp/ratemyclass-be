package services

import "go.uber.org/fx"

var Module = fx.Module("services", fx.Provide(
	fx.Annotate(
		NewAuthServiceImpl, fx.As(new(AuthService)),
	),
	fx.Annotate(
		NewProfessorService, fx.As(new(ProfessorService)),
	),
	fx.Annotate(
		NewCourseRatingService, fx.As(new(CourseRatingService)),
	),
	fx.Annotate(
		NewCourseService, fx.As(new(CourseService)),
	),
	fx.Annotate(
		NewJwtService, fx.As(new(JwtService)),
	),
	fx.Annotate(
		NewSchoolService, fx.As(new(SchoolService)),
	),
))
