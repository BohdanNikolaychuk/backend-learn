package dto

import "time"

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

type CreateUserResponse struct {
	ID        string    `json:"id"`
	Version   int64     `json:"version"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUserRequest struct {
	FullName string `json:"full_name"`
}

type ChangePasswordUserRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
