package model

type User struct {
	FirstName string `form:"first_name" json:"first_name" binding:"required"`
	LastName  string `form:"last_name" json:"last_name" binding:"required"`
	Gender    string `form:"gender" json:"gender" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
}
