package requests

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `json:"email" binding:"required"`
	Name  string `json:"name" binding:"required"`
}
