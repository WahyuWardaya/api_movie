package routes

import (
	"api_movie/controllers"

	"github.com/gin-gonic/gin"
)

func MovieRoutes(router *gin.Engine) {
	movies := router.Group("/movies")
	{
		movies.GET("", controllers.GetMovies)
		movies.POST("", controllers.CreateMovies)
		movies.GET("/:id", controllers.GetMoviesByID)
		movies.PUT("/:id", controllers.UpdateMovies)
		movies.DELETE("/:id", controllers.DeleteMovies)
	}
}
