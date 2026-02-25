package routes

import (
	"sharedup/app/src/middlewares"
	"sharedup/app/src/posts/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func RoutesPost(router *gin.Engine) {
	routes := router.Group("/api/v1/post")

	create := dependencies.GetCreatePostController().Create
	update := dependencies.GetUpdatePostController().Update
	delete := dependencies.GetDeletePostController().Delete
	getAll := dependencies.GetGetAllPostController().GetAll
	getByUser := dependencies.GetGetPostByUserController().GetByUser

	like := dependencies.GetIncrementLikeController().Increment
	dislike := dependencies.GetIncrementDislikeController().Increment

	routes.POST("/",middlewares.AuthMiddleware(),create)
	routes.PUT("/:id", middlewares.AuthMiddleware(), update)
	routes.DELETE("/:id", middlewares.AuthMiddleware(), delete)
	routes.GET("/", middlewares.AuthMiddleware(), getAll)
	routes.GET("/me", middlewares.AuthMiddleware(), getByUser)

	routes.PUT("/:id/like", middlewares.AuthMiddleware(), like)
	routes.PUT("/:id/dislike", middlewares.AuthMiddleware(), dislike)
	
	wsHandler := dependencies.GetWebSocketHandler()
	router.GET("/ws/posts", wsHandler.Handle)
}