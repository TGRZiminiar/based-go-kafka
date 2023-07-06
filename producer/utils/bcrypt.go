package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	temp := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(temp, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CheckPassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
