package controllers

import (
	"net/http"
	"website/server/models"

	"github.com/gin-gonic/gin"
)

// Store all of the handle functions
// Get product
func GetALLProductsFromDB(c *gin.Context) {
	resp := models.GetProductALL("funko")
	sCode := http.StatusOK

	if resp.ErrMsg != nil {
		sCode = http.StatusNotFound
	}

	c.JSON(sCode, gin.H{
		"data": resp,
	})
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")

	resp := models.GetProductByID("funko", id)
	sCode := http.StatusOK

	if resp.ErrMsg != nil {
		sCode = http.StatusNotFound
	}

	c.JSON(sCode, gin.H{
		"data": resp,
	})
}
