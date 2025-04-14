package main

import (
	"api_movie/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	router := gin.Default()
	router.Run(":3000")
}