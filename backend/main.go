package main


import (
	"os"
	"fmt"	
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/DylanCoon99/portfolio/backend/database"
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
		// GET projects

		// GET pictures

		// 


	}
	//protected := r.Group("/api/admin")
	// Need to create a JWT autheticator for admin panel
	// include special admin apis for adding new projects and pictures


	r.Run()

}


