package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       int8      `json:"age"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Lastname string `json:"lastname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Age      *int8  `json:"age" binding:"required,max=90"`
}
