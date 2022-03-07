package Interface

import (
	"bookms/Database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BorrowRet struct {
	Success bool `json:"success"`
}

func BorrowBook(c *gin.Context) {
	bookId, _ := strconv.Atoi(c.PostForm("bookId"))

	claim, err := VertifyToken(c)
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: "false", Msg: err.Error()})
		return
	}

	authName := claim.Name
	authPass := claim.Password

	err = Database.Borrow(Database.Book{ID: bookId}, Database.User{Username: authName, Password: authPass})
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: "false", Msg: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, BorrowRet{Success: true})
	return
}

func ReturnBook(c *gin.Context) {

	bookId, _ := strconv.Atoi(c.PostForm("bookId"))

	claim, err := VertifyToken(c)
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: "false", Msg: err.Error()})
		return
	}

	authName := claim.Name
	authPass := claim.Password

	err = Database.Return(Database.Book{ID: bookId}, Database.User{Username: authName, Password: authPass})
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: "false", Msg: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, BorrowRet{Success: true})
	return
}
