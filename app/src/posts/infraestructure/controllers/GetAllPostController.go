package controllers

import (
	"net/http"
	usecases "sharedup/app/src/posts/application/useCases"

	"github.com/gin-gonic/gin"
)

type GetAllPostController struct {
	useCase usecases.GetAllPostUseCase
}

func NewGetAllPostController(useCase *usecases.GetAllPostUseCase) *GetAllPostController {
	return &GetAllPostController{useCase: *useCase}
}

func (c *GetAllPostController) GetAll(ctx *gin.Context) {

	posts, err := c.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, posts)
}