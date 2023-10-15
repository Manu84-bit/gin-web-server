package main

import (
	"github.com/gin-gonic/gin"
)


func main(){
	//crear router
	router:= gin.Default()

	//get request
	router.GET("/gin-greeting", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, stranger!",
		})
	})

	//run server
	router.Run()

}