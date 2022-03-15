package user

import "time"

type UserModel struct {
	ID        int       `json:"id"`
	Firstname string    `json:"first_name" validate:"required"`
	Lastname  string    `json:"last_name" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	RoleID    int       `json:"role_id"`
	Active    bool      `json:"active"`
	ImageName string    `json:"image_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLoginModel struct {
	ID       int    `json:"id"`
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	RoleID   int    `json:"role_id"`
	Active   bool   `json:"active"`
}

type ResponseModel struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type UserProfileModel struct {
	Firstname string `json:"first_name" validate:"required"`
	Lastname  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	RoleID    int    `json:"role_id" validate:"required"`
	ImageName string `json:"image_name"`
}
