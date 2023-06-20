package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/prince-konwea/restuarant-management-api/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controller.GetFoods())
	incomingRoutes.GET("/foods/:food_id", controller.GetFood())
	incomingRoutes.POST("/foods", controller.CreateFood())
	incomingRoutes.GET("/foods/:food_id", controller.UpdateFood())
}
