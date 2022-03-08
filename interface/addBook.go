package Interface

import (
	"bookms/Database"
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AddBookRet struct {
	Success bool `json:"success"`
}

func AddBook(c *gin.Context) {
	title := c.PostForm("title")
	author := c.PostForm("author")
	num, _ := strconv.Atoi(c.PostForm("num"))
	publisher := c.PostForm("publisher")
	publicDate := c.PostForm("publicDate")
	libName := c.PostForm("libName")
	libLocation := c.PostForm("libLocation")
	//checkout := c.PostForm("checkout") == "true"

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
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}

	idh := sha256.Sum256([]byte(title + author + publisher + publicDate + libName + libLocation))

	var ids string

	for _, value := range idh {
		ids = ids + string(value)
	}
	id, _ := strconv.Atoi(ids)
	Database.AddBook(Database.Book{
		ID:               id,
		Title:            title,
		CopyNum:          num,
		Author:           author,
		PublicationDate:  publicDate,
		PublisherName:    publisher,
		PublisherAddress: "",
	})

	for i := 1; i <= num; i++ {
		Database.AddCopy(Database.Copy{
			BookID:          id,
			CopyID:          id + i,
			LibraryName:     libName,
			LibraryLocation: libLocation,
		})

	}

}
