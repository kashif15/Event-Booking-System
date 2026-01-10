package auth

import (
	"errors"
	"event-booking-api/pkg/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId int64, role string) (string, error) {

	claims := jwt.MapClaims{
		"user_id" : userId,
		"role" : role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret :=  config.Get("JWT_SECRET")

	if secret == "" {
		return "", errors.New("JWT_SECRET not set")
	}

	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	secret :=  config.Get("JWT_SECRET")

	if secret == "" {
		return nil, errors.New("JWT_SECRET not set")
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil

}