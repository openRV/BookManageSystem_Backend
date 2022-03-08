package main

import (
	"bookms/Interface"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
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
	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
