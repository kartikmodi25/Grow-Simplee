package main

import (
	"backend-assignment/config"
	"backend-assignment/database/postgres"
	"backend-assignment/server"
	"context"
	"flag"

	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()
	configPath := flag.String("config", ".", "path to directory containing config.yaml, default current directory")
	flag.Parse()
	err := config.Init(ctx, *configPath, config.ConfigName, "local")
	if err != nil {
		log.Err(err).Msg("failed to initialise the config")
		return
	}
	var c config.Database
	c.Host = "localhost"
	c.Name = "postgres3"
	c.Password = "252900"
	c.Port = 5432
	c.Username = "postgres"
	db, err := postgres.New(c)
	if err != nil {
		log.Err(err).Msg("failed to connect to database, exiting")
	}
	err = db.AutoMigrate()
	if err != nil {
		log.Err(err).Msg("failed to create tables in database")
		return
	}
	server, err := server.Init(ctx, db)
	if err != nil {
		log.Err(err).Msg("Service.Init failed: , exiting")
		return
	}
	err = server.Start(ctx)
	if err != nil {
		log.Err(err).Msg("Service.Run failed: , exiting")
		return
	}
}
