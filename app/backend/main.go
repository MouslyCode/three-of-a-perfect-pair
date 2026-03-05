package main

import (
	"github.com/MouslyCode/three-of-a-perfect-pair/backend/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error Load .env !")
	}

	database.Connect()
}
