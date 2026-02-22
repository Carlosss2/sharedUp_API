package main

import (
	"sharedup/app/src/helpers"
	registerDependencies "sharedup/app/src/register/infraestructure/dependencies"
	registerRoutes"sharedup/app/src/register/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main(){

	registerDependencies.InitDependencies()

	r := gin.Default()
	helpers.InitCORS(r)

	registerRoutes.Routes(r)

	r.Run(":8081")

}