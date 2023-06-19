package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/prince-konwea/restuarant-management-api/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controller.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemByOrder())
	incomingRoutes.POST("/orderItem", controller.CreateOrderItem())
	incomingRoutes.PATCH("/orderItem/orderItem_id", controller.UpdateOrderItem())
}
