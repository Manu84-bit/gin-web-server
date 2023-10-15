package main

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)


func main(){	

	p:= Product{
		"MacBook Pro", 500.00, true,
	}

	jsonData, err := json.Marshal(p)
	if err!= nil{
		log.Fatal(err)
	}

	//crear router
	router:= gin.Default()

	//get request
	router.GET("/gin-greeting", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, stranger!",
			"product": jsonData,
		})
	})

	//run server
	router.Run()

}