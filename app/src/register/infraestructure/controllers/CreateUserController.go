package controllers

import (
	usecases "sharedup/app/src/register/application/useCases"
	"sharedup/app/src/register/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	useCase *usecases.CreateUserUseCase
	
}

func NewCreateUserController(useCase *usecases.CreateUserUseCase)*CreateUserController{
	return &CreateUserController{useCase: useCase}
}

func (controller *CreateUserController) Execute(ctx *gin.Context){

	var user entities.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
	ctx.JSON(400, gin.H{
		"error": "JSON inválido",
	})
	return
}

if err := entities.ValidateUser(user); err != nil {
	ctx.JSON(400, gin.H{
		"error": err.Error(),
	})
	return
}

if err := controller.useCase.Execute(user); err != nil {
	ctx.JSON(500, gin.H{
		"error": err.Error(),
	})
	return
}
ctx.JSON(201, gin.H{
	"message": "Usuario creado exitosamente",
})

}
