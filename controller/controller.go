package controller

import (
	"example-auth/model"
	"example-auth/model/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome !"})
}

func ValidateData(c *gin.Context) {
	var userReg model.User

	if err := c.ShouldBindJSON(&userReg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "validated!"})
}

func Login(c *gin.Context) {
	var userReg model.User

	if err := c.ShouldBindJSON(&userReg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isValid := service.ValidateUser(userReg)
	if isValid {
		c.JSON(http.StatusOK, gin.H{"message": "validated!"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "You shall not pass!"})
	}
}

func RegisterUser(c *gin.Context) {
	var userReg model.User

	if err := c.ShouldBindJSON(&userReg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.NewUser(userReg)
	c.JSON(http.StatusOK, gin.H{"message": "done!"})
}
