package main

import (
	"github.com/gin-gonic/gin"

	"bookms/Interface"
)

func main() {
	router := gin.Default()
	router.POST("/login", Interface.PostUser)
	router.POST("/register", Interface.RegisterUser)
	router.GET("/faculty/users", Interface.GetUsers)
	router.GET("/admin/users", Interface.GetUsers)
	router.Run("localhost:8080")
}
