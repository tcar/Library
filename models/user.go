package models

type User struct {
	ID    int
	Email string
	Name  string
}

type PrivateUserDetails struct {
	ID       int
	Password string
	Salt     string
}
