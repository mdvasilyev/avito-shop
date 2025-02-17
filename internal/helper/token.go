package helper

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

func NewToken(userID int) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := &jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(30 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

func ValidateToken(tokenString string) (*jwt.MapClaims, error) {
	secret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return &claims, nil
}
