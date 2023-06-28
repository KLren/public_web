package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"website/server/models"

	"github.com/gin-gonic/gin"
)

// Store all of the handle functions
// Get product
func GetALLUsersFromDB(c *gin.Context) {
	resp := models.GetUserALL("users")
	sCode := http.StatusOK

	if resp.ErrMsg != nil {
		sCode = http.StatusNotFound
	}

	c.JSON(sCode, gin.H{
		"data": resp,
	})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	resp := models.GetUserByID("users", id)
	sCode := http.StatusOK

	if resp.ErrMsg != nil {
		sCode = http.StatusNotFound
	}

	c.JSON(sCode, gin.H{
		"data": resp,
	})
}

func LoginUser(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		return
	}

	dataMap := make(map[string]string)
	json.Unmarshal(jsonData, &dataMap)
	resp := models.GetUserByEmailPassword(dataMap)
	sCode := http.StatusOK

	if resp.ErrMsg != nil {
		sCode = http.StatusNotFound
	}

	c.JSON(sCode, gin.H{
		"data": resp,
	})
}
