package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	prod "github.com/Manu84-bit/gin-web-server/product"
	"github.com/gin-gonic/gin"
)


func main(){	

	p:= prod.Product{
		Name: "MacBook Pro", Price: 500.00, Published: true,
	}

	//Read json file and parse to a Product slice
	var products []prod.Product
	jsonFile, err1 := os.ReadFile("../products.json")
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

	//Path params:
	router.GET("/productos/:id", func(ctx *gin.Context){
			producto, ok := findProductsById(products,ctx.Param("id"))
		if ok=="found"{
			ctx.String(200, "Nombre: %v; id %v", producto.Name, ctx.Param("id"))
		} else {
			ctx.String(400, "Producto no encontrado.")
		}
	})
   
	//Query params:
	router.GET("/productos", func(ctx *gin.Context) {
		producto, ok :=findProductsByName(products, ctx.Query("nombre"))
			if ok=="found"{
			ctx.String(200, "Nombre: %v; precio: %v", producto.Name, producto.Price)
		} else {
			ctx.String(400, "Producto no encontrado.")
		}
	})

	//Query params to create product:
	router.GET("/productparams", func(ctx *gin.Context) {
		newId, err1 := strconv.Atoi(ctx.Query("id"))
		newName:= ctx.Query("nombre")
		newPrice, err3 := strconv.ParseFloat(ctx.Query("precio"), 64)
		newStock, err4 := strconv.Atoi(ctx.Query("stock"))
		newCode:= ctx.Query("código")
		newDate:= ctx.Query("vencimiento")
		newPublished, er7 := strconv.ParseBool(ctx.Query("publicado"))
		errors := []error{err1,err3, err4, er7}
		for _, e := range errors{
			if e!=nil{
				ctx.String(400, "Error: wrong data product.")
			}
			break
		}

		newProduct:= prod.Product{Id: newId, Name: newName, Price: newPrice, Stock: newStock, Code: newCode, Expiration: newDate, Published: newPublished}
		ctx.JSON(http.StatusOK, gin.H{
			"newProduct": newProduct,
		})
		products = append(products, newProduct)
		
	})

	//Query params to search by quantity:
router.GET("/searchbyquantity", func(ctx *gin.Context) {
		min, err1 := strconv.Atoi(ctx.Query("min"))
		max, err2 := strconv.Atoi(ctx.Query("max"))
		if err1 != nil || err2 != nil{
			ctx.String(400, "Error: wrong data")
		} else {
			searchedProducts, ok := findProductsByQuantity(products, min, max)
			if ok=="found"{
			ctx.JSON(200, gin.H{
				"foundProducts": searchedProducts,
			})
		} else {
			ctx.String(400, "No existen productos en ese rango de cantidades.")
		}
		}
})


router.GET("/buy", func(ctx *gin.Context) {
		name:=ctx.Query("name")
		code:=ctx.Query("code")
		units, err:= strconv.Atoi(ctx.Query("units"))
		
		if err != nil{
			ctx.String(400, "Error: wrong data")
		} else {
			product, ok := findProductsByName(products, name)
			if ok!="" && product.Name == name {
					detail:= Detalle{
							code, name, units, product.Price * float64(units),
						}
			
			ctx.JSON(200, gin.H{
				"detalleCompra": detail,
			})
		}else {
			ctx.String(400, "La información del producto es incorrecta.")
		}
	}
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



func findProductsById(products []prod.Product, id string) (prod.Product, string){
	p:= prod.Product {
	}
	ok := ""
	for _, product := range products {
		if fmt.Sprint(product.Id) == id {
			p= product
			break
		}
	}
	if p.Id != 0 {
		ok = "found"
	}
return p, ok
}

func findProductsByName(products []prod.Product, name string) (prod.Product, string){
	p:= prod.Product {
	}
	ok := ""
	for _, product := range products {
		if fmt.Sprint(product.Name) == name {
			p= product
			break
		}
	}
	if p.Name != "" {
		ok = "found"
	}
return p, ok
}

func findProductsByQuantity(products []prod.Product, min, max int) ([]prod.Product, string){
	searchedProducts:= []prod.Product {
	}
	ok := ""
	for _, product := range products {
		if product.Stock >= min && product.Stock <= max {
			searchedProducts = append(searchedProducts, product)
		}
	}
	if len(searchedProducts) > 0 {
		ok = "found"
	}
return searchedProducts, ok
}

type Detalle struct {
	Code string
	Name string
	Units int
	Total float64
}