package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	prod "github.com/Manu84-bit/gin-web-server/product"
	"github.com/gin-gonic/gin"
)


func main(){	

	p:= prod.Product{
		Name: "MacBook Pro", Price: 500.00, Published: true,
	}

	//Read json file and parse to a Product slice
	var products []prod.Product
	jsonFile, err1 := os.ReadFile("products.json")
	if err1!= nil {
		log.Fatal(err1)
	}
	json.Unmarshal(jsonFile, &products)


	//Format Product slice to print as a string with json format
	data, err := json.Marshal(products)
		if err!= nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))



	//crear router
	router:= gin.Default()

	//get request
	router.GET("/gin-greeting", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, stranger!",
			"product": p,
			"products": products,
		})
	})

	//Grupo de endpoints:
	gopher := router.Group("/perfil")
	{
		gopher.GET("/foto", func(ctx *gin.Context){
			ctx.JSON(200, gin.H{
				"message":"foto here",
			})
		})
		gopher.GET("/datos",func(ctx *gin.Context){
			ctx.JSON(200, gin.H{
				"message":"data here",
			})
		})
	}

	//run server
	router.Run()

}