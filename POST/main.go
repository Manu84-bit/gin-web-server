package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username" binding:"required`
	Password string `json:"password" binding:"required`
}

func main() {
	server := gin.Default()

	server.POST("/login", func(ctx *gin.Context) {
		var login Login 
		ctx.BindJSON(&login)
		ctx.JSON(http.StatusOK, gin.H{
			"status": login.Username,
		})
	})


	server.Run(":3000")
}