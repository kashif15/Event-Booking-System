package main

import (
	"fmt"

	"event-booking-api/internal/auth"
	"event-booking-api/pkg/config"
)

func main() {

	config.Load()
	token, err := auth.GenerateToken(1, "USER")
	fmt.Println("Token:", token, err)

	claims, err := auth.ValidateToken(token)
	fmt.Println("Claims:", claims, err)
}
