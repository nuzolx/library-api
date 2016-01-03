package models

import "time"

type Book struct {
	Id				int				`json:"id"`
	Name			string      	`json:"name"`
	Author	 		string       	`json:"author"`
	ReleaseDate		time.Time		`json:"releaseDate"`
}