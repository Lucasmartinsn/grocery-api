package Supplier

import (
	"strconv"

	models "github.com/Lucasmartinsn/grocery-api/Models/Supplier"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
		var data models.Supplier
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding json",
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
		var data models.Product
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding json",
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
		var data models.Batch
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding json",
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
		var supplier models.Supplier
		err = c.ShouldBindJSON(&supplier)
		if err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding json",
			})
			return
		}
		row, err := models.UpdatedSupplier(id, supplier)
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
	} else if params["product"] {
		var product models.Product
		err = c.ShouldBindJSON(&product)
		if err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding json",
			})
			return
		}
		valueBool, _ := strconv.ParseBool(c.Query("volume"))
		row, err := models.UpdatedProduct(id, product,valueBool)
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
	} else if params["batch"] {
		var batch models.Batch
		err = c.ShouldBindJSON(&batch)
		if err != nil {
			c.JSON(400, gin.H{
				"Error":   err.Error(),
				"Message": "error decoding json",
			})
			return
		}
		row, err := models.UpdatedBatch(id, batch)
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
