package Employee

import (
	"strconv"

	models "github.com/Lucasmartinsn/grocery-api/Models/Employee"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func EmployeeCreate(c *gin.Context) {
	var employee models.Employee
	err := c.ShouldBindJSON(&employee)
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   err.Error(),
			"Mensage": "error decoding json",
		})
		return
	}

	err = models.CreationEmployee(employee)
	if err != nil {
		c.JSON(400, gin.H{
			"Error":    err.Error(),
			"menssage": "error when trying to insert",
		})

	} else {
		c.JSON(201, gin.H{
			"menssage": "registered successfully",
		})
	}
}

func EmployeeUpdate(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   err.Error(),
			"Mensage": "error when parsing the id",
		})
		return
	}
	keys := []string{"pass", "name", "office", "active", "admin"}
	params := make(map[string]bool)
	for _, key := range keys {
		params[key] = false
	}

	for _, key := range keys {
		valueStr := c.Query(key)
		if valueStr != "" {
			// Converter string para bool
			valueBool, err := strconv.ParseBool(valueStr)
			if err != nil {
				continue
			}
			params[key] = valueBool
		}
	}

	var employee models.Employee
	err = c.ShouldBindJSON(&employee)
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   err.Error(),
			"Mensage": "error decoding json",
		})
		return
	}

	row, err := models.UpdateEmployee(id, params, employee)
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   err.Error(),
			"Mensage": "error updating register",
		})
		return
	}

	if row == 0 {
		c.JSON(500, gin.H{
			"Error": "internal database error",
		})
		return
	} else if row == 404 {
		c.JSON(500, gin.H{
			"Error": "method not found",
		})
		return
	} else if row > 1 {
		c.JSON(400, gin.H{
			"Error": "multiple updated records",
		})
	} else {
		c.JSON(200, gin.H{
			"Mensage": "registration updated successfully",
		})
	}
}

func EmployeeList(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   err.Error(),
			"Mensage": "error when parsing the id",
		})
		return
	}
	status, err := strconv.ParseBool(c.Query("status"))
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "error when converting string to bool",
		})
	}

	resp, err := models.SearchEmployees(id, status)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"employee": resp,
		})
	}
}

func EmployeeDelete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   err.Error(),
			"Mensage": "error when parsing the id",
		})
		return
	}

	row, err := models.DeleteEmployee(id)
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   err.Error(),
			"Mensage": "error deleted register",
		})
		return
	}
	if row == 0 {
		c.JSON(500, gin.H{
			"Error": "internal database error",
		})
		return
	} else if row == 404 {
		c.JSON(500, gin.H{
			"Error": "method not found",
		})
		return
	} else if row > 1 {
		c.JSON(400, gin.H{
			"Error": "multiple deleted records",
		})
	} else {
		c.JSON(200, gin.H{
			"Mensage": "registration deleted successfully",
		})
	}
}