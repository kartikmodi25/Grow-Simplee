package main

import (
	"backend-assignment/database/models"
	"backend-assignment/database/postgres"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	db, err := postgres.GetConnection()
	db.AutoMigrate(&models.Movie{})
	db.AutoMigrate(&models.User{})
	fmt.Println(db, err)
}
