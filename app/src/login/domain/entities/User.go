package entities

import "github.com/go-playground/validator/v10"


type User struct{
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// Modelo INTERNO para login (NO se expone)
type UserWithPassword struct {
	Id           int    `gorm:"column:iduser"`     // Verifica si tu PK se llama iduser
	Name         string `gorm:"column:name"`
	Email        string `gorm:"column:email"`
	PasswordHash string `gorm:"column:passwordHash"` // Debe ser EXACTO a la BD
	Career 		 string `gorm:"column:career"`
}

// DTO de salida
type UserResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Career   string  `json:"career"`
}

func ValidateUser(user User) error {
	validate := validator.New()
	return validate.Struct(user)
}
