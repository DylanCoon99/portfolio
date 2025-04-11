package controllers

import (
	//"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/DylanCoon99/portfolio/backend/database"
	//"github.com/DylanCoon99/portfolio/backend/models"
)




func GetAllProjects(c *gin.Context) {

	projects, err := database.GetAllProjects()
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database."})
	}

	c.JSON(http.StatusOK, projects)

}


func GetProjectByID(c *gin.Context) {

	project_id := c.Param("project_id")

	project, err := database.GetProjectByID(project_id)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database."})
	}

	c.JSON(http.StatusOK, project)

}



func GetAllImages(c *gin.Context) {


	c.JSON(http.StatusOK, gin.H{"resp": "to be implemented"})

}