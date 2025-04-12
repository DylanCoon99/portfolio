package controllers

import (
	//"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/DylanCoon99/portfolio/backend/database"
	"github.com/DylanCoon99/portfolio/backend/models"
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


func CreateProject(c *gin.Context) {

	var newProject models.Project

	if err := c.ShouldBindJSON(&newProject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input", "details": err.Error()})
		return
	}

	if err := database.CreateProject(&newProject); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newProject)

}


func GetAllImages(c *gin.Context) {


	c.JSON(http.StatusOK, gin.H{"resp": "to be implemented"})

}



func TestAuth(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"resp": "to be implemented"})

}