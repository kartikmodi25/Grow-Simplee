package main

import (
	"backend-assignment/database/models"
	"backend-assignment/database/postgres"
	"backend-assignment/router"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World!")
	db, _ := postgres.GetConnection()
	db.AutoMigrate(&models.Movie{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.UserToken{})
	r := router.SetupRouter(db)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal("Invalid port number")
	}
	if err := r.Run(":" + strconv.Itoa(port)); err != nil {
		log.Fatal(err)
	}
}
