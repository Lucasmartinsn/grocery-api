package Routers

import (
	"github.com/gin-gonic/gin"

	customer "github.com/Lucasmartinsn/grocery-api/Handles/Customer"
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
			Login.POST("customer/login", customer.LoginCustomer)
			Login.POST("employee/login/customer", customer.LoginEmployee)
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
		Customer := main.Group("customer", Middlewares.Auth_customer())
		{
			Customer.GET("/", customer.CustomerList)
			Customer.GET("/address/:id", customer.CustomerListAddress)
			Customer.GET("/card/:id", customer.CustomerListCard)
			Customer.POST("/", customer.CustomerCreate)
			Customer.PUT("/:id", customer.CustomerUpdate)
			Customer.DELETE("/:id", customer.CustomerDelete)

		}
	}
	return router
}
