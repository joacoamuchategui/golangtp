package main

import (
	"github.com/gin-gonic/gin"
	"tp-golang/service"
)
func main() {
	service := service.NewUserService()
	router := gin.Default()
	router.GET("/users", service.GetUser)
	router.POST("/users", service.AddUser)
	router.Run(":8080")
}