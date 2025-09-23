package crypto

import (
	"Food_Delivery_Management/utils"

	"golang.org/x/crypto/bcrypt"
)

func BcryptHash(password string) ([]byte, error) {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	utils.IsNotNilError(err, "bcrypt", "bcrypt is issue")
	return hashpassword, nil
}

func BcryptCompare(hashpassword []byte, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hashpassword, []byte(password))
	return err == nil, err
}
