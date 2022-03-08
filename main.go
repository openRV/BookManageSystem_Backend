package main

import (
	"bookms/Interface"
	"bookms/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.Use(middleware.Cors())

	router.POST("/login", Interface.PostUser)
	router.POST("/register", Interface.RegisterUser)
	router.GET("/faculty/users", Interface.GetUsers)
	router.GET("/admin/users", Interface.GetUsers)
	router.GET("/users/account", Interface.GetAccount)
	router.POST("/users/modify", Interface.ModifyAccount)
	router.DELETE("/admin/delusers", Interface.DeleteUser)
	router.DELETE("/faculty/delusers", Interface.DeleteUser)
	router.GET("/users/books", Interface.SearchBook)
	router.DELETE("/faculty/delbook", Interface.DeleteBook)
	router.POST("/users/borrow", Interface.BorrowBook)
	router.POST("/users/return", Interface.ReturnBook)
	router.GET("/users/borrowing", Interface.Borrowing)
	router.GET("/users/borrowed", Interface.Borrowed)
	router.POST("faculty/addbook", Interface.AddBook)
	router.POST("admin/addbook", Interface.AddBook)
	router.DELETE("/faculty/delpaper", Interface.DeletePaper)
	router.GET("/users/conference", Interface.SearchConferencePaper)
	router.GET("/users/journal", Interface.SearchJournalPaper)

	router.GET("/test", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Test OK")
	})

	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
