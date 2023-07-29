package router

import (
	middlewares "backend-assignment/middleware"
	"backend-assignment/server"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	// Initialize Gin router
	r := gin.Default()

	// Register routes
	r.POST("/users", server.RegisterUser(db))
	r.POST("/login", server.LoginUser(db))
	r.POST("/rate", middlewares.JwtMiddleware, server.RateMovie(db))
	r.GET("/listmovies", server.ListMovies(db))
	r.GET("/listmovieratings", server.ListMovieRatings(db))
	return r
}
