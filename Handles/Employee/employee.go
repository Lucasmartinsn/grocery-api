package Employee

import (
	"encoding/json"
	"strconv"

	key "github.com/Lucasmartinsn/grocery-api/Configs/confEnv"
	models "github.com/Lucasmartinsn/grocery-api/Models/Employee"
	descrypt "github.com/Lucasmartinsn/grocery-api/Services/EncryptionResponse"
	Services "github.com/Lucasmartinsn/grocery-api/Services/EncryptionToken"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var requestData struct {
	Data string `json:"data"`
}

func EmployeeCreate(c *gin.Context) {
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
	var employee models.Employee
	if err = json.Unmarshal([]byte(decryptedData), &employee); err != nil {
		c.JSON(400, gin.H{
			"Error":   err.Error(),
			"Message": "error decoding decrypted json",
		})
		return
	}
	if err = models.CreationEmployee(employee); err != nil {
		c.JSON(400, gin.H{
			"Error":   err.Error(),
			"Message": "error when trying to insert",
		})

	} else {
		c.JSON(201, gin.H{
			"Message": "registered successfully",
		})
	}
}
func EmployeeUpdate(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   err.Error(),
			"Message": "error when parsing the id",
		})
		return
	}
	keys := []string{"all", "pass", "name", "office", "active", "admin"}
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
	var employee models.Employee
	if err = json.Unmarshal([]byte(decryptedData), &employee); err != nil {
		c.JSON(400, gin.H{
			"Error":   err.Error(),
			"Message": "error decoding decrypted json",
		})
		return
	}

	if row, err := models.UpdateEmployee(id, params, employee); err != nil {
		c.JSON(400, gin.H{
			"Error":   err.Error(),
			"Message": "error updating register",
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
			"Error": "multiple updated records",
		})
	} else {
		c.JSON(200, gin.H{
			"Message": "registration updated successfully",
		})
	}
}
func EmployeeList(c *gin.Context) {
	resp, err := models.SearchEmployees(c.Query("id"), c.Query("status"))
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
			"Message": "error when parsing the id",
		})
		return
	}

	row, err := models.DeleteEmployee(id)
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
func ValidateLogin(c *gin.Context) {
	var login models.Employee
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   err.Error(),
			"Message": "error when decoding json",
		})
		return
	}

	register, err := models.Validate(login.Cpf, login.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
		return
	}

	token, err := Services.NewJWTService_Default().GenerateToken_Default(register.Id)
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
