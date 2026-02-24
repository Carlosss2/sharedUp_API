package routes

import (
	"sharedup/app/src/login/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine){
	routes := router.Group("/api/v1/login")
	{
		loginController := dependencies.GetLogInController()
		routes.POST("/", loginController.Execute)
	}
}