package handlers

import (
	"AustemServer/models"
	"AustemServer/posgtresql/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllProducts(c *gin.Context) {
	var products []models.Product
	db.DB.Find(&products)

	c.JSON(http.StatusOK, products)
}

func PostProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	//correct json
	//{
	//	"id":"5277485f-bf5f-44b1-8d43-684c9964475c",
	//	"name":"juj",
	//	"count":22,
	//	"cost":15,
	//	"isSelected": false
	//}

	if err := db.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create product"})
		return
	}

	c.JSON(http.StatusOK, product)
}
