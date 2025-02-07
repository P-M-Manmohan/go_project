package controller

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(tokenString string) error {

    secretKey := []byte(GoDotEnvVariable("SECRET_KEY"))
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      return secretKey, nil
    })
  
    if err != nil {
      return err
    }

  
    if !token.Valid {
      return fmt.Errorf("invalid token")
    }
  
   return nil
}

func GenerateTocken(username string) (string, error) {
	secretKey := []byte(GoDotEnvVariable("SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tockenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tockenString, nil
}
