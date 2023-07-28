package server

import (
	"backend-assignment/database/postgres"
	"backend-assignment/requests"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		user := requests.User{}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse request body"})
			return
		}
		exist, err := postgres.CheckExistingUser(db, user.Email)
		if err != nil {
			log.Error().Err(err).Str("email", user.Email).Msg("error checking user in db")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create new user, try again"})
			return
		}
		if exist {
			log.Error().Err(err).Str("email", user.Email).Msg("user already exists in db")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user already exist"})
			return
		}
		password := "qwerty"
		accessToken := "qwerty"
		err = postgres.CreateUser(db, user.Name, user.Email, password, accessToken)
		c.JSON(http.StatusCreated, user)
	}
}
