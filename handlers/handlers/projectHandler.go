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

func GetSingleProject(c *gin.Context) {
	var project models.Project
	if err := db.DB.Where("name = ?", project.Name).First(&project).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "project": project.Project})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "project": project.Project})

}
func AddProject(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"stasus": "bad", "error": "invalid json"})
		return
	}
	if err := db.DB.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"stasus": "bad", "error": "failed to add"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"stasus": "good", "message": "project created successfully"})
}
