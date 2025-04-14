package routes

import (
	"api_movie/controllers"

	"github.com/gin-gonic/gin"
)

func RoleRoutes(router *gin.Engine) {
	users := router.Group("/roles")
	{
		users.GET("", controllers.GetUsers)
		users.POST("", controllers.CreateUser)
		users.GET("/:id", controllers.GetUserByID)
		users.PUT("/:id", controllers.UpdateUsers)
		users.DELETE("/:id", controllers.DeleteUsers)
	}
}