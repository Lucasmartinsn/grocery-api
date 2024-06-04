package Customer

import (
	"strconv"

	Models "github.com/Lucasmartinsn/grocery-api/Models/Customer"
	Employee "github.com/Lucasmartinsn/grocery-api/Models/Employee"
	Services "github.com/Lucasmartinsn/grocery-api/Services/EncryptionToken"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func LoginCustomer(c *gin.Context) {
	var login Models.Customer
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   err.Error(),
			"Message": "error when decoding json",
		})
		return
	}

	register, err := Models.ValidateCustomer(login.Cpf, login.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
		return
	}

	token, err := Services.NewJWTService_Customer().GenerateToken_Customer(register.Id)
	if err != nil {
		c.JSON(500, gin.H{
			"Error":   err.Error(),
			"Message": "token not generated! return null data",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"token": token,
		})
	}
}
func LoginEmployee(c *gin.Context) {
	var login Employee.Employee
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   err.Error(),
			"Message": "error when decoding json",
		})
		return
	}

	register, err := Employee.Validate(login.Cpf, login.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
		return
	}

	token, err := Services.NewJWTService_Customer().GenerateToken_Customer(register.Id)
	if err != nil {
		c.JSON(500, gin.H{
			"Error":   err.Error(),
			"Message": "token not generated! return null data",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"token": token,
		})
	}
}
func CustomerCreate(c *gin.Context) {
	keys := []string{"customer", "address", "card"}
	params := make(map[string]bool)
	for _, key := range keys {
		params[key] = false
	}
	for _, key := range keys {
		valueStr := c.Query(key)
		if valueStr != "" {
			// convert string to bool
			valueBool, err := strconv.ParseBool(valueStr)
			if err != nil {
				continue
			}
			params[key] = valueBool
		}
	}
	if params["customer"] {
		var data Models.Customer
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding json",
			})
			return
		}
		id, err := Models.CreationCustomer(data)
		if err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error when trying to insert",
			})
			return
		} else {
			c.JSON(201, gin.H{
				"Message": "registered successfully",
				"Id":      id,
			})
			return
		}
	} else if params["address"] {
		var data Models.Address
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding json",
			})
			return
		}
		if err := Models.CreationAddress(data); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error when trying to insert",
			})
			return
		} else {
			c.JSON(201, gin.H{
				"Message": "registered successfully",
			})
			return
		}
	} else if params["card"] {
		var data Models.Credit_card
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding json",
			})
			return
		}
		if err := Models.CreationCard(data); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error when trying to insert",
			})
			return
		} else {
			c.JSON(201, gin.H{
				"Message": "registered successfully",
			})
			return
		}
	} else {
		c.JSON(404, gin.H{
			"message": "method not found",
		})
		return
	}
}
func CustomerList(c *gin.Context) {
	resp, err := Models.SearchCustomer(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"customer": resp,
		})
	}
}
func CustomerListAddress(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
		return
	}
	resp, err := Models.SearchCustomer_address(id)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
		return
	} else {
		c.JSON(200, resp)
	}
}
func CustomerListCard(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
		return
	}
	resp, err := Models.SearchCustomer_card(id)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
		return
	} else {
		c.JSON(200, resp)
	}
}
func CustomerUpdate(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   err.Error(),
			"Message": "error when parsing the id",
		})
		return
	}
	keys := []string{"customer", "pass", "address", "card"}
	params := make(map[string]bool)
	for _, key := range keys {
		params[key] = false
	}

	for _, key := range keys {
		valueStr := c.Query(key)
		if valueStr != "" {
			// convert string to bool
			valueBool, err := strconv.ParseBool(valueStr)
			if err != nil {
				continue
			}
			params[key] = valueBool
		}
	}
	if params["customer"] {
		var customer Models.Customer
		if err = c.ShouldBindJSON(&customer); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding json",
			})
			return
		}
		row, err := Models.UpdatedCustomer(id, customer, params["pass"])
		if err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error updating register",
			})
			return
		}
		if row != 1 {
			c.JSON(500, gin.H{
				"Error": "internal database error",
			})
			return
		} else {
			c.JSON(200, gin.H{
				"Message": "registration updated successfully",
			})
			return
		}
	} else if params["address"] {
		var address Models.Address
		if err = c.ShouldBindJSON(&address); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding json",
			})
			return
		}
		row, err := Models.UpdatedCustomer_address(id, address)
		if err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error updating register",
			})
			return
		}
		if row != 1 {
			c.JSON(500, gin.H{
				"Error": "internal database error",
			})
			return
		} else {
			c.JSON(200, gin.H{
				"Message": "registration updated successfully",
			})
			return
		}
	} else if params["card"] {
		var card Models.Credit_card
		if err = c.ShouldBindJSON(&card); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding json",
			})
			return
		}
		row, err := Models.UpdatedCustomer_card(id, card)
		if err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error updating register",
			})
			return
		}
		if row != 1 {
			c.JSON(500, gin.H{
				"Error": "internal database error",
			})
			return
		} else {
			c.JSON(200, gin.H{
				"Message": "registration updated successfully",
			})
			return
		}
	} else {
		c.JSON(404, gin.H{
			"message": "method not found",
		})
		return
	}
}
func CustomerDelete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   err.Error(),
			"Message": "error when parsing the id",
		})
		return
	}
	row, err := Models.DeleteCustumer(id)
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   err.Error(),
			"Message": "error deleted register",
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
			"Message": "registration deleted successfully",
		})
	}
}