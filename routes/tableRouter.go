package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/prince-konwea/restuarant-management-api/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controller.GetTables())
	incomingRoutes.GET("/tables/:table_id", controller.GetTable())
	incomingRoutes.POST("/tables", controller.CreateTable())
	incomingRoutes.PATCH("/invoices/invoice_id", controller.UpdateTable())
}
