package controllers

import (
	"net/http"
	"strconv"
	usecases "sharedup/app/src/posts/application/useCases"

	"github.com/gin-gonic/gin"
)

type DeletePostController struct {
	useCase usecases.DeletePostUseCase
}

func NewDeletePostController(useCase *usecases.DeletePostUseCase) *DeletePostController {
	return &DeletePostController{useCase: *useCase}
}

func (c *DeletePostController) Delete(ctx *gin.Context) {
	// Obtener el ID del post desde la URL
	idParam := ctx.Param("id")
	postID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de post inválido"})
		return
	}

	// Obtener el ID del usuario desde el middleware
	iduser_str, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	userID := iduser_str.(int)

	if err := c.useCase.Execute(postID, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Publicacion eliminada"})
}