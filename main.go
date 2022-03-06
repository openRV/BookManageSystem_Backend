package main

import (
	"github.com/gin-gonic/gin"

	"bookms/Interface"
)

func main() {
	router := gin.Default()
	router.POST("/login", Interface.PostUser)
	router.POST("/register", Interface.RegisterUser)
	router.Run("localhost:8080")
}
