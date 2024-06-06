package Customer

import (
	"encoding/json"
	"strconv"

	key "github.com/Lucasmartinsn/grocery-api/Configs/confEnv"
	Models "github.com/Lucasmartinsn/grocery-api/Models/Customer"
	Employee "github.com/Lucasmartinsn/grocery-api/Models/Employee"
	descrypt "github.com/Lucasmartinsn/grocery-api/Services/EncryptionResponse"
	Services "github.com/Lucasmartinsn/grocery-api/Services/EncryptionToken"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var requestData struct {
	Data string `json:"data"`
}

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
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding json",
			})
			return
		}
		decryptedData, err := descrypt.DecryptData(requestData.Data, []byte(key.Variable()))
		if err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decrypting data",
			})
			return
		}
		var data Models.Customer
		if err = json.Unmarshal([]byte(decryptedData), &data); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding decrypted json",
			})
			return
		}
		if id, err := Models.CreationCustomer(data); err != nil {
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
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding json",
			})
			return
		}
		decryptedData, err := descrypt.DecryptData(requestData.Data, []byte(key.Variable()))
		if err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decrypting data",
			})
			return
		}
		var data Models.Address
		if err = json.Unmarshal([]byte(decryptedData), &data); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding decrypted json",
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
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding json",
			})
			return
		}
		decryptedData, err := descrypt.DecryptData(requestData.Data, []byte(key.Variable()))
		if err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decrypting data",
			})
			return
		}
		var data Models.Credit_card
		if err = json.Unmarshal([]byte(decryptedData), &data); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding decrypted json",
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
	if id, err := uuid.Parse(c.Param("id")); err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
		return
	} else if resp, err := Models.SearchCustomer_address(id); err != nil {
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
	if resp, err := Models.SearchCustomer_card(id); err != nil {
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
		if err = c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding json",
			})
			return
		}
		decryptedData, err := descrypt.DecryptData(requestData.Data, []byte(key.Variable()))
		if err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decrypting data",
			})
			return
		}
		var customer Models.Customer
		if err = json.Unmarshal([]byte(decryptedData), &customer); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding decrypted json",
			})
			return
		}
		if row, err := Models.UpdatedCustomer(id, customer, params["pass"]); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error updating register",
			})
			return
		} else if row != 1 {
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
		if err = c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding json",
			})
			return
		}
		decryptedData, err := descrypt.DecryptData(requestData.Data, []byte(key.Variable()))
		if err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decrypting data",
			})
			return
		}
		var address Models.Address
		if err = json.Unmarshal([]byte(decryptedData), &address); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding decrypted json",
			})
			return
		}

		if row, err := Models.UpdatedCustomer_address(id, address); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error updating register",
			})
			return
		} else if row != 1 {
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
		if err = c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding json",
			})
			return
		}
		decryptedData, err := descrypt.DecryptData(requestData.Data, []byte(key.Variable()))
		if err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decrypting data",
			})
			return
		}
		var card Models.Credit_card
		if err = json.Unmarshal([]byte(decryptedData), &card); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding decrypted json",
			})
			return
		}
		if row, err := Models.UpdatedCustomer_card(id, card); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error updating register",
			})
			return
		} else if row != 1 {
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
	if row, err := Models.DeleteCustumer(id); err != nil {
		c.JSON(400, gin.H{
			"Error":   err.Error(),
			"Message": "error deleted register",
		})
		return
	} else if row == 0 {
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
