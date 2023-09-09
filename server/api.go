package server

import (
	middlewares "backend-assignment/middleware"

	"github.com/gin-gonic/gin"
)

func addV1Api(r *gin.Engine, s *Server) {
	authRoutes := r.Group("/auth")
	// Register routes
	authRoutes.POST("/register", RegisterUser(s))
	authRoutes.POST("/login", LoginUser(s))
	r.POST("/rate", middlewares.JwtMiddleware, RateMovie(s))
	r.GET("/listmovies", ListMovies(s))
	r.GET("/listmovieratings", ListMovieRatings(s))
}
