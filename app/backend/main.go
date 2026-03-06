package main

import (
	"log"

	"github.com/MouslyCode/three-of-a-perfect-pair/backend/database"
	"github.com/MouslyCode/three-of-a-perfect-pair/backend/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error Load .env !")
	}

	database.Connect()

	r := gin.Default()
	router.Routes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}

}
