package main

import (
	"api_movie/config"
	"api_movie/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	router := gin.Default()
	routes.UserRoutes(router)
	routes.RoleRoutes(router)
	routes.GenreRoutes(router)
	routes.ActorRoutes(router)
	routes.DirectorRoutes(router)

	router.Run(":3000")
}