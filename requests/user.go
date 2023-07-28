package requests

type User struct {
	Email string `json:"email" binding:"required"`
	Name  string `json:"name" binding:"required"`
}

type LoginUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Movie struct {
	Name   string `json:"name" binding:"required"`
	Rating int8   `json:"rating" binding:"required"`
}
