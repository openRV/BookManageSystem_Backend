package Interface

import (
	"bookms/Database"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DelBookRet struct {
	Success bool `json:"success"`
}

func DeleteBook(c *gin.Context) {

	bookId, _ := strconv.Atoi(c.Param("bookId"))

	claim, err := VertifyToken(c)
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}

	authName := claim.Name
	authPass := claim.Password

	property, err := Database.GetUserProperty(Database.User{Username: authName, Password: authPass})
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}

	if property == Database.Student {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: "your property is student, not able to delete book"})
		return
	}

	err = Database.DelBook(Database.Book{ID: bookId})
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}

	c.IndentedJSON(http.StatusOK, DelBookRet{Success: true})
}
