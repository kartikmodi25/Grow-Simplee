package main

import (
	"backend-assignment/database/postgres"
	"backend-assignment/router"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	db, err := postgres.GetConnection()
	if err != nil {
		log.Err(err).Msg("database connection failed, exiting")
		return
	}
	err = db.AutoMigrate()
	if err != nil {
		log.Err(err).Msg("failed to create tables in database")
		return
	}
	r := router.SetupRouter(db)
	err = godotenv.Load()
	if err != nil {
		log.Err(err).Msg("error loading .env file")
		return
	}
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Err(err).Msg("invalid port number")
		return
	}
	if err := r.Run(":" + strconv.Itoa(port)); err != nil {
		log.Err(err)
		return
	}
}
