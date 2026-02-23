package routes

import (
	"sharedup/app/src/middlewares"
	"sharedup/app/src/posts/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func RoutesPost(router *gin.Engine) {
	routes := router.Group("/api/v1/post")

	create := dependencies.GetCreatePostController().Create

	routes.POST("/",middlewares.AuthMiddleware(),create)
}