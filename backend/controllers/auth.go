package controllers


import (
	"os"
	"log"
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
	//"github.com/joho/godotenv"
	//"golang.org/x/crypto/bcrypt"
	"github.com/DylanCoon99/portfolio/backend/utils"
)



type LoginInput struct {
	Username string `json"username" binding:"required"`
	Password string `json"password" binding:"required"`
}




func VerifyUsernameAndPassword(passwordSupplied, validPassword, usernameSupplied, validUsername string) bool {
	return (passwordSupplied == validPassword) && (usernameSupplied == validUsername)
}



func Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	// We need to verify login
	token, err := VerifyLogin(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or Password is incorrect", "error_message": err.Error()})
		return 
	}

	c.JSON(http.StatusOK, gin.H{"token":token})

}




func VerifyLogin(input LoginInput) (string, error) {

	/*err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
		return "", err
	}
	*/

	ADMIN_USER := os.Getenv("ADMIN_USER")
	ADMIN_PW   := os.Getenv("ADMIN_PW")


	log.Printf("DEBUG: %s %s\n", ADMIN_USER, ADMIN_PW)

	verify := VerifyUsernameAndPassword(input.Password, ADMIN_PW, input.Username, ADMIN_USER)

	if !verify {
		return "", errors.New("Username or Password is incorrect")
	}

	token, err := utils.GenerateToken(input.Username)

	if err != nil {
		return "", err
	}

	return token, nil

}