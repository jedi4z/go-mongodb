package handler

import "github.com/gin-gonic/gin"

type RequestHandler interface {
	UserHandler
	HandlePing(c *gin.Context)
}

type UserHandler interface {
	HandleNewUser(c *gin.Context)
	HandleListUsers(c *gin.Context)
	HandleGetUser(c *gin.Context)
	HandleUpdateUser(c *gin.Context)
}