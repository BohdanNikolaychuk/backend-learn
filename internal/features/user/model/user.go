package model

import "time"

type User struct {
	ID        string     `json:"id"`
	Version   int64      `json:"version"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	FullName  string     `json:"full_name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
