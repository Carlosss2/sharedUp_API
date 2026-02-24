package controllers

import (
	"sharedup/app/src/helpers"
	usecases "sharedup/app/src/login/application/useCases"
	"sharedup/app/src/login/domain/entities"

	"github.com/gin-gonic/gin"
)

type LogInController struct {
	usecase *usecases.LogInUseCase
}

func NewLogInController(usecase *usecases.LogInUseCase) *LogInController {
	return &LogInController{usecase: usecase}
}

func (controller *LogInController) Execute(ctx *gin.Context) {
	var user entities.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": "JSON inválido"})
		return
	}

	if err := entities.ValidateUser(user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userResponse, err := controller.usecase.Execute(user)
	if err != nil {
		ctx.JSON(401, gin.H{"error": "Credenciales inválidas"})
		return
	}

	token, err := helpers.GenerateJWT(userResponse.Id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error al generar token"})
		return
	}

	ctx.JSON(200, gin.H{
		"data":  userResponse,
		"token": token,
	})
}

