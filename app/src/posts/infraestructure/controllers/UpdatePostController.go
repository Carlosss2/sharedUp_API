package controllers

import (
	"net/http"
	"strconv"
	usecases "sharedup/app/src/posts/application/useCases"
	"sharedup/app/src/posts/domain/entities"

	"github.com/gin-gonic/gin"
)

type UpdatePostController struct {
	useCase usecases.UpdatePostUseCase
}

func NewUpdatePostController(useCase *usecases.UpdatePostUseCase) *UpdatePostController {
	return &UpdatePostController{useCase: *useCase}
}

func (c *UpdatePostController) Update(ctx *gin.Context) {
	var post entities.Post
	
	idParam := ctx.Param("id")
	postID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de post inválido"})
		return
	}
	post.Id = postID

	iduser_str, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	post.IdUser = iduser_str.(int)

	// Bindear el JSON para obtener el nuevo título y texto
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.useCase.Execute(post); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Publicacion actualizada"})
}