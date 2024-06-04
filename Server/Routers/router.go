package Routers

import (
	"github.com/gin-gonic/gin"

	employee "github.com/Lucasmartinsn/grocery-api/Handles/Employee"
	supllier "github.com/Lucasmartinsn/grocery-api/Handles/Supplier"
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
		Employee := main.Group("employee", Middlewares.Auth_default())
		{
			Employee.GET("/", employee.EmployeeList)
			Employee.POST("/", employee.EmployeeCreate)
			Employee.PUT("/:id", employee.EmployeeUpdate)
			Employee.DELETE("/:id", employee.EmployeeDelete)
		}
		Supplier := main.Group("supplier", Middlewares.Auth_default())
		{
			Supplier.GET("/", supllier.SupplierList)
			Supplier.GET("/product/:id", supllier.SupplierListProduct)
			Supplier.GET("/batch/:id", supllier.SupplierListBatch)
			Supplier.POST("/", supllier.SupplierCreate)
			Supplier.PUT("/:id", supllier.SupplierUpdate)
		}
	}
	return router
}
