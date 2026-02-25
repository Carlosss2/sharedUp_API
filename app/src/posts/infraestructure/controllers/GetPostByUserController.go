package controllers

import (
	"net/http"
	usecases "sharedup/app/src/posts/application/useCases"

	"github.com/gin-gonic/gin"
)

type GetPostByUserController struct {
	useCase usecases.GetPostByUserUseCase
}

func NewGetPostByUserController(useCase *usecases.GetPostByUserUseCase) *GetPostByUserController {
	return &GetPostByUserController{useCase: *useCase}
}

func (c *GetPostByUserController) GetByUser(ctx *gin.Context) {

	idUserStr, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	idUser, ok := idUserStr.(int)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el ID del usuario"})
		return
	}

	posts, err := c.useCase.Execute(idUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, posts)
}