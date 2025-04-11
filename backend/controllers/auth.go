package controllers


import (
	"golang.org/x/crypto/bcrypt"
	"github.com/DylanCoon99/portfolio/backend/utils/token"
)




func EncryptPassword(u *models.User) error {

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.PasswordHash = string(passwordHashed)

	u.Name = html.EscapeString(strings.TrimSpace(u.Name))

	return nil
}
