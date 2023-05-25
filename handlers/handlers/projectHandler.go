package handlers

import (
	"AustemServer/models"
	"AustemServer/posgtresql/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllProjects(c *gin.Context) {
	var projects []models.Project
	db.DB.Find(&projects)

	c.JSON(http.StatusOK, projects)
}
