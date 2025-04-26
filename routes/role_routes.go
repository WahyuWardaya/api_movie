package routes

import (
	"api_movie/controllers"

	"github.com/gin-gonic/gin"
)

func RoleRoutes(router *gin.Engine) {
	users := router.Group("/roles")
	{
		users.GET("", controllers.GetRoles)
		users.POST("", controllers.CreateRoles)
		users.GET("/:id", controllers.GetRolesByID)
		users.PUT("/:id", controllers.UpdateRoles)
		users.DELETE("/:id", controllers.DeleteRoles)
	}
}