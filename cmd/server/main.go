package main

import (
	"example-auth/controller"
	"example-auth/repository"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	err := repository.Connect()
	if err != nil {
		fmt.Println("Error creating database")
	}
	r := gin.Default()

	r.GET("/", controller.Index)
	r.POST("/register", controller.RegisterUser)
	r.GET("/login", controller.Login)
	r.GET("/test", controller.ValidateData)

	r.Run(":8080")
	repository.Close()
}
