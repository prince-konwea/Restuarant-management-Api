package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/prince-konwea/restuarant-management-api/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/user_id", controller.GetUser())
	incomingRoutes.POST("/users/SignUp", controller.SignUp())
	incomingRoutes.POST("/users/Login", controller.Login())
}
