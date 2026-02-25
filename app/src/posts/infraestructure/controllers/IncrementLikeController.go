package controllers

import (
	"net/http"
	"strconv"
	usecases "sharedup/app/src/posts/application/useCases"

	"github.com/gin-gonic/gin"
)

type IncrementLikeController struct {
	useCase usecases.IncrementLikeUseCase
}

func NewIncrementLikeController(useCase *usecases.IncrementLikeUseCase) *IncrementLikeController {
	return &IncrementLikeController{useCase: *useCase}
}

func (c *IncrementLikeController) Increment(ctx *gin.Context) {

	idParam := ctx.Param("id")
	postID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.useCase.Execute(postID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Like incrementado"})
}