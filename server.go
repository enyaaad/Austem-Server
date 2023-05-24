package server

import (
	"AustemServer/handlers/handlers"
	"github.com/gin-gonic/gin"
)

func StartAPI() {

	r := gin.Default()

	productGroup := r.Group("/products")
	{
		productGroup.GET("/getall", handlers.GetAllProducts)
		productGroup.POST("/add", handlers.PostProduct)

	}
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
