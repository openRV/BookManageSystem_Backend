package main

import (
	"github.com/gin-gonic/gin"

	"bookms/Interface"
)

func main() {
	router := gin.Default()
	router.POST("/user", Interface.PostUser)
	router.Run("localhost:8080")
}
