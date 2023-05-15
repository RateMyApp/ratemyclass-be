package services

type RegisterCommand struct {
	Email     string
	Password  string
	Firstname string
	Lastname  string
}

type LoginCommand struct {
	Email string
	Password string
}

type UserDetails struct {
	Firstname    string
	Lastname     string
	Email        string
}