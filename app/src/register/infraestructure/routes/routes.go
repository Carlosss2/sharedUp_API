package routes

import (
	"sharedup/app/src/register/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine){
	routes := router.Group("/api/v1/register")
	{
		createUserController := dependencies.GetCreateUserController()
		routes.POST("/", createUserController.Execute)
	}

}