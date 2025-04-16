package controllers

import (
	"fmt"
	"os"
	"net/http"
	"net/smtp"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"github.com/DylanCoon99/portfolio/backend/database"
	"github.com/DylanCoon99/portfolio/backend/models"
)


type ContactForm struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}



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



func Contact(c *gin.Context) {


	var newContactForm ContactForm

	if err := c.ShouldBindJSON(&newContactForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input", "details": err.Error()})
		return
	}

	// Send email
	name    := newContactForm.Name
	email   := newContactForm.Email
	message := newContactForm.Message

	myEmail := "dylcoon99@gmail.com"
	password := os.Getenv("EMAIL_PASSWORD")
	smtpHost := "smtp.gmail.com"
    smtpPort := "587"
    body := fmt.Sprintf("Name: %s\nEmail: %s\nMessage:\n%s", name, email, message)

	auth := smtp.PlainAuth("", myEmail, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, myEmail, []string{myEmail}, []byte("Subject: New Portfolio Message\r\n\r\n" + body))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email", "details": err.Error()} )
		return
	}

	c.JSON(http.StatusCreated, newContactForm)
	return
}


func UploadImage(c *gin.Context) {

	file, err := c.FormFile("file")

	project_id := c.Param("project_id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file found."})
		return
	}

	path := filepath.Join("static/images", file.Filename)

	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image."})
		return
	}

	// update the path in the database
	if err := database.UploadImage(project_id, file.Filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update image path in database."})
		return
	}


	c.JSON(http.StatusOK, gin.H{"path": "/images" + file.Filename})

	return

}


