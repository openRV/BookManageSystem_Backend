package main

import (
	"github.com/gin-gonic/gin"

	"bookms/Interface"
)

func main() {
	router := gin.Default()
	router.GET("/user", Interface.GetUser)
	router.Run("localhost:8080")
}
