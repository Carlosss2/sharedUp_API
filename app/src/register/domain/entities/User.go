package entities

import "github.com/go-playground/validator/v10"

type User struct {
    Id       int    `json:"id"`
    Name     string `json:"name" validate:"required"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8"`
	Career string  `json:"career" validate:"required,career"`
}

//Mandar a llamar en el controller
func ValidateUser(user User) error {
	validate := validator.New()
	return validate.Struct(user)
}