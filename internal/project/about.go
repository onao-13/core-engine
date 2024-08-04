package project

import "time"

type About struct {
	Name        string    `json:"name"`
	Version     string    `json:"version"`
	Author      string    `json:"author"`
	Email       string    `json:"email"`
	Description string    `json:"description"`
	DateCreated time.Time `json:"dateCreated"`
}
