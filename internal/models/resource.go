package models

import "time"

type Resource struct{
	ID int `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Role string  `json:"role" db:"role"`
	Stack string 	`json:"stack" db:"stack"`
	Email string  `json:"email" db:"email"`
	Status string	`json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdateAt time.Time `json:"update_at" db:"updateAt"`
	
}