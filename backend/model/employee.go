
package model

import "time"

type Employee struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Department  string    `json:"department"`
	Designation string    `json:"designation"`
	JoiningDate time.Time `json:"joining_date"`
	Status      string    `json:"status"`
}
