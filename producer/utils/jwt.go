package utils

import (
	"errors"
	"io/ioutil"
	"log"

	"github.com/golang-jwt/jwt/v4"
)

type CookieCliam struct {
	UserId   string
	UserName string
	Email    string
	Role     string
}

func (c CookieCliam) Validator(data CookieCliam) error {

	if len(data.UserId) < 5 {
		return errors.New("userid is missing")

	} else if data.Email == "" {
		return errors.New("email is required")

	} else if data.UserName == "" {
		return errors.New("username is required")

	} else if data.Role == "" {
		return errors.New("role can't be empty")

	} else {
		return nil
	}
}

func GenToken(data CookieCliam) (tokenString string, err error) {

	err = data.Validator(data)
	if err != nil {
		return "", err
	}

	privateKeyBytes, err := ioutil.ReadFile("private.pem")
	if err != nil {
		return "", err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = data.UserId
	claims["userName"] = data.UserName
	claims["email"] = data.Email
	claims["role"] = data.Role

	tokenString, err = token.SignedString(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}

func GetDataFromToken(tokenString string) (CookieCliam, error) {

	privateKeyBytes, err := ioutil.ReadFile("private.pem")
	if err != nil {
		return CookieCliam{}, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(privateKeyBytes)
	if err != nil {
		return CookieCliam{}, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		return CookieCliam{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return CookieCliam{}, errors.New("invalid token claims")
	}

	userId, ok := claims["userId"].(string)
	if !ok {
		return CookieCliam{}, errors.New("invalid userId claim")
	}

	userName, ok := claims["userName"].(string)
	if !ok {
		return CookieCliam{}, errors.New("invalid userName claim")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return CookieCliam{}, errors.New("invalid email claim")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return CookieCliam{}, errors.New("invalid role claim")
	}

	return CookieCliam{
		UserId:   userId,
		UserName: userName,
		Email:    email,
		Role:     role,
	}, nil
}
