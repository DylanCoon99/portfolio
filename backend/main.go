package main


import (
	"os"
	"fmt"	
	"log"
	//"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/DylanCoon99/portfolio/backend/database"
	"github.com/DylanCoon99/portfolio/backend/controllers"
)


func main() {


	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}


	DbHost := os.Getenv("DB_HOST")
	//fmt.Printf("This is the host name: %v", DbHost)
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")


	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require", DbHost, DbUser,  DbPassword, DbName, DbPort)
	err = database.ConnectDB(dsn)

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	// Define our routes here
	api := r.Group("/api")
	{
		// GET all projects
		api.GET("/projects", controllers.GetAllProjects)
		// GET project by id
		api.GET("/projects/:project_id", controllers.GetProjectByID)
		// GET pictures
		api.GET("/images", controllers.GetAllImages)
		// 


	}
	//protected := r.Group("/api/admin")
	// Need to create a JWT autheticator for admin panel
	// include special admin apis for adding new projects and pictures


	r.Run()

}


