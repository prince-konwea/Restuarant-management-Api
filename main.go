package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prince-konwea/restuarant-management-api/database"
	"github.com/prince-konwea/restuarant-management-api/middleware"
	"github.com/prince-konwea/restuarant-management-api/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.openCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoicesRoutes(router)

	router.Run(":" + port)

}
