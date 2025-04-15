package routes

import (
	"api_movie/controllers"

	"github.com/gin-gonic/gin"
)

func ActorRoutes(router *gin.Engine) {
	actors := router.Group("/actors")
	{
		actors.GET("", controllers.GetActors)
		actors.POST("", controllers.CreateActors)
		actors.GET("/:id", controllers.GetActorsByID)
		actors.PUT("/:id", controllers.UpdateActors)
		actors.DELETE("/:id", controllers.DeleteActors)
	}
}