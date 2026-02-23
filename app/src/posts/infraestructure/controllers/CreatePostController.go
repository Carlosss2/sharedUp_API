package controllers

import (
	"net/http"
	usecases "sharedup/app/src/posts/application/useCases"
	"sharedup/app/src/posts/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreatePostController struct {
	useCase usecases.CreatePostUseCase
}

func NewCreatePostController(useCase *usecases.CreatePostUseCase)*CreatePostController{
	return &CreatePostController{useCase: *useCase}
}

func(c *CreatePostController) Create(ctx *gin.Context){
	var post entities.Post
	iduser_str,exist := ctx.Get("user_id")
	 if !exist {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
        return
    }

	iduser, ok := iduser_str.(int)
    if !ok {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el ID del usuario"})
        return
    }
	post.IdUser = iduser

	  if err := ctx.ShouldBindJSON(&post); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    if err := c.useCase.Execute(post); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, gin.H{"message": "Publicacion registrada"})
    
}