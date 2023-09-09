package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func router(ctx context.Context, s *Server) *gin.Engine {
	router := gin.New()
	router.GET("/ping", ping)
	addV1Api(router, s)
	return router
}
