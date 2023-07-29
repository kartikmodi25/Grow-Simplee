package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Movie struct {
	gorm.Model
	Name   string  `json:"name"`
	Rating float64 `json:"rating"`
	Count  int64   `json:"Count"`
}
