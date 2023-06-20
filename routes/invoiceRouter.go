package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/prince-konwea/restuarant-management-api/controllers"
)

func InvoicesRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoice", controller.GetInvoices())
	incomingRoutes.GET("/invoice/:invoice_id", controller.GetInvoice())
	incomingRoutes.POST("/invoices", controller.CreateInvoice())
	incomingRoutes.PATCH("/invoices/invoice_id", controller.UpdateInvoice())
}
