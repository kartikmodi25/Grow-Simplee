package router

import (
	"backend-assignment/server"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	// Initialize Gin router
	r := gin.Default()

	// Register routes
	r.POST("/users", server.RegisterUser(db))

	return r
}
