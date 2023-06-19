package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prince-konwea/restuarant-management-api/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/user_id", controllers.GetUser())
	incomingRoutes.POST("/users/SignUp", controllers.SignUp())
	incomingRoutes.POST("/users/Login", controllers.Login())
}
