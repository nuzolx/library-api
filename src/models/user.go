package models

type User struct {
	Login		string      	`json:"login"`
	Password 	string       	`json:"password"`
	Email 		string			`json:"email"`
}