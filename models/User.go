package models

type User struct {
	Name string
	Email string
	Username string
	Password string
	Projects []Project
}