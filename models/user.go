package models

import "time"

type User struct {
	ID         string
	Email      string
	Department string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
