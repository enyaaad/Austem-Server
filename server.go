package server

import (
	handlers2 "AustemServer/handlers"
	"AustemServer/handlers/handlers"
	"github.com/gin-gonic/gin"
)

func StartAPI() {

	r := gin.Default()

	r.Use(handlers2.CORSMiddleware())

	productGroup := r.Group("/products")
	{
		productGroup.GET("/getall", handlers.GetAllProducts)
		productGroup.POST("/add", handlers.PostProduct)
	}
	userGroup := r.Group("/user")
	{
		userGroup.POST("/auth", handlers.Autharization)
	}
	projectGroup := r.Group("/project")
	{
		projectGroup.GET("/getall", handlers.GetAllProjects)
		projectGroup.GET("/get", handlers.GetSingleProject)
		projectGroup.POST("/add", handlers.AddProject)
	}

	err := r.Run(":8080")

	if err != nil {
		return
	}
}
