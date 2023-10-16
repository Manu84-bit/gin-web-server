package main

import (
	// 	"encoding/json"
	// 	"log"
	// 	"os"
	"fmt"
	"log"
	"net/http"
	"strconv"

	emp "github.com/Manu84-bit/gin-web-server/employee"
	"github.com/gin-gonic/gin"
)


func main(){
	employees := []emp.Employee {
		{Id: 1, Nombre: "Juan", Activo: true},
		{Id: 2, Nombre: "María", Activo: false},
		{Id: 3, Nombre: "Carlos", Activo: true},
		{Id: 4, Nombre: "Manu", Activo: false},
	}

	server := gin.Default()
	server.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Welcome, Stranger")
	})
	//GET to optain the list of employees:
	server.GET("/api/employees", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"employees": employees,
		})
	})

	//Path param using "id":
	server.GET("/api/employees/:id", func(ctx *gin.Context) {
		employee, ok := findEmployeeById(employees, ctx.Param("id"))
		if ok != "" {
			ctx.JSON(200, gin.H{
			"employee": employee,
		})
		} else {
			ctx.String(400, "Employee not found")
		}
	})

	//Query param using "active":
	server.GET("api/employees/find", func(ctx *gin.Context) {
		employeesA := findActiveEmployees(employees, ctx.Query("active"))
		ctx.JSON(200, gin.H{
			"employeesFound": employeesA, 
		})
	})

	//Query param to create employee:
	//Crear una ruta /employeesparams que nos permita crear un empleado a través de los params
	// y lo devuelva en formato JSON.
	server.GET("/api/employees/create", func(c *gin.Context) {
		
		id, err1 := strconv.Atoi(c.Query("id"))
		name := c.Query("name") 
		active, err2 := strconv.ParseBool(c.Query("active"))

		if err1 == nil && err2 == nil {
			c.JSON(http.StatusOK, gin.H{
					"id":id,
					"name": name,
					"active": active,
				})

		employees = append(employees, emp.Employee{Id: id, Nombre: name, Activo: active })	
		}else {
			c.String(400,"Error: check the data introduced.")
			log.Fatal(err1)
			log.Fatal(err2)
		}
	})

	server.Run()
}

func findEmployeeById(employees []emp.Employee, id string) (emp.Employee, string){
	employee := emp.Employee {	
	}
	ok := ""
	for _, e := range employees {
		if fmt.Sprint(e.Id) == id {
			employee = e
			ok = "found"
			break
		}
	}
	return employee, ok
}

func findActiveEmployees(employees []emp.Employee, active string) ([]emp.Employee){
	activeEmployees := []emp.Employee {	
	}
	inactiveEmployees := []emp.Employee {	
	}
	for _, e := range employees {
		if fmt.Sprint(e.Activo) == "true" {
			activeEmployees = append(activeEmployees, e)
		} else if fmt.Sprint(e.Activo)== "false" {
			inactiveEmployees = append(inactiveEmployees, e)
		}
	}
	if active == "true" {
		return activeEmployees
	}else if active == "false" {
		return inactiveEmployees
	} else {
		return []emp.Employee{}
	}
}
