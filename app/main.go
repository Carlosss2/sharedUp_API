package main

import (
	"sharedup/app/src/helpers"
	logInDependencies "sharedup/app/src/login/infraestructure/dependencies"
	logInRoutes "sharedup/app/src/login/infraestructure/routes"
	postDependencies "sharedup/app/src/posts/infraestructure/dependencies"
	postRoutes "sharedup/app/src/posts/infraestructure/routes"
	registerDependencies "sharedup/app/src/register/infraestructure/dependencies"
	registerRoutes "sharedup/app/src/register/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main(){

	registerDependencies.InitDependencies()
	logInDependencies.InitDependencies()
	postDependencies.InitPost()

	r := gin.Default()
	helpers.InitCORS(r)

	registerRoutes.Routes(r)
	logInRoutes.Routes(r)
	postRoutes.RoutesPost(r)

	r.Run(":8081")

}