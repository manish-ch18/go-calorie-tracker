package main

import (
	"go-calorie-tracker/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := "8000"
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/entry/create", routes.AddEntry)
	router.GET("/entries", routes.GetEntries)
	router.GET("/entry/:id/", routes.GetEntryByID)
	router.GET("/ingredients/:ingredient/", routes.GetEntriesByIngredients)
	router.PUT("/entry/update/:id", routes.UpdateEntry)
	router.PUT("/ingredients/update/:id", routes.UpdateIngredients)
	router.DELETE("/entry/delete/:id", routes.DeleteEntry)

	router.Run(":" + port)
}
