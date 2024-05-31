package Routers

import (
	"github.com/gin-gonic/gin"

	employee "github.com/Lucasmartinsn/grocery-api/Handles/Employee"
	"github.com/Lucasmartinsn/grocery-api/Server/Middlewares"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	router.Use(Middlewares.CORSMiddleware())
	router.Use(gin.Recovery())
	main := router.Group("api")
	{
		Login := main.Group("")
		{
			Login.POST("employee/login", employee.ValidateLogin)
		}
		Employee := main.Group("employee", Middlewares.Auth())
		{
			Employee.GET("/", employee.EmployeeList)
			Employee.POST("/", employee.EmployeeCreate)
			Employee.PUT("/:id", employee.EmployeeUpdate)
			Employee.DELETE("/:id", employee.EmployeeDelete)
		}
	}
	return router
}
