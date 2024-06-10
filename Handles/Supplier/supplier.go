package Supplier

import (
	"encoding/json"
	"strconv"

	key "github.com/Lucasmartinsn/grocery-api/Configs/confEnv"
	models "github.com/Lucasmartinsn/grocery-api/Models/Supplier"
	descrypt "github.com/Lucasmartinsn/grocery-api/Services/EncryptionResponse"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var requestData struct {
	Data string `json:"data"`
}

func descrypts(c *gin.Context) (string, error) {
	if err := c.ShouldBindJSON(&requestData); err != nil {
		return "", err
	}
	decryptedData, err := descrypt.DecryptData(requestData.Data, []byte(key.Variable()))
	if err != nil {
		return "", err
	}
	return decryptedData, nil
}
func SupplierCreate(c *gin.Context) {
	keys := []string{"supplier", "product", "batch"}
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
	if params["supplier"] {
		decryptedData, err := descrypts(c)
		if err != nil || decryptedData == "" {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decrypting data",
			})
			return
		}
		var data models.Supplier
		if err = json.Unmarshal([]byte(decryptedData), &data); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding decrypted json",
			})
			return
		}
		if err := models.CreationSupplier(data); err != nil {
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
	} else if params["product"] {
		decryptedData, err := descrypts(c)
		if err != nil || decryptedData == "" {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decrypting data",
			})
			return
		}
		var data models.Product
		if err = json.Unmarshal([]byte(decryptedData), &data); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding decrypted json",
			})
			return
		}
		if err := models.CreationProduct(data); err != nil {
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
	} else if params["batch"] {
		decryptedData, err := descrypts(c)
		if err != nil || decryptedData == "" {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decrypting data",
			})
			return
		}
		var data models.Batch
		if err = json.Unmarshal([]byte(decryptedData), &data); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding decrypted json",
			})
			return
		}
		if err := models.CreationBatch(data); err != nil {
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
func SupplierList(c *gin.Context) {
	resp, err := models.SearchSupplier()
	if err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"supplier": resp,
		})
	}
}
func SupplierListProduct(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
		return
	}
	resp, err := models.SearchSupplier_product(id)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
		return
	} else {
		c.JSON(200, resp)
	}
}
func SupplierListBatch(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
		return
	}
	resp, err := models.SearchSupplier_bacth(id)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": err.Error(),
		})
		return
	} else {
		c.JSON(200, resp)
	}
}
func SupplierUpdate(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   err.Error(),
			"Message": "error when parsing the id",
		})
		return
	}
	keys := []string{"supplier", "product", "batch"}
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
	if params["supplier"] {
		decryptedData, err := descrypts(c)
		if err != nil || decryptedData == "" {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decrypting data",
			})
			return
		}
		var supplier models.Supplier
		if err = json.Unmarshal([]byte(decryptedData), &supplier); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding decrypted json",
			})
			return
		}
		if row, err := models.UpdatedSupplier(id, supplier); err != nil {
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
	} else if params["product"] {
		decryptedData, err := descrypts(c)
		if err != nil || decryptedData == "" {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decrypting data",
			})
			return
		}
		var product models.Product
		if err = json.Unmarshal([]byte(decryptedData), &product); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding decrypted json",
			})
			return
		}
		valueBool, _ := strconv.ParseBool(c.Query("volume"))
		if row, err := models.UpdatedProduct(id, product, valueBool); err != nil {
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
	} else if params["batch"] {
		decryptedData, err := descrypts(c)
		if err != nil || decryptedData == "" {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decrypting data",
			})
			return
		}
		var batch models.Batch
		if err = json.Unmarshal([]byte(decryptedData), &batch); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding decrypted json",
			})
			return
		}
		if row, err := models.UpdatedBatch(id, batch); err != nil {
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
