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

	json := make(map[string]interface{})
	c.BindJSON(&json)

	title := json["title"].(string)
	author := json["author"].(string)
	num, _ := strconv.Atoi(json["num"].(string))
	publisher := json["publisher"].(string)
	publicDate := json["publicDate"].(string)
	libName := json["libName"].(string)
	libLocation := json["libLocation"].(string)
	//checkout := c.PostForm("checkout") == "true"

	publicDate = publicDate[:10]

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

	var id int
	for _, value := range idh {
		id = id + int(value)
	}
	Database.AddBook(Database.Book{
		ID:               id,
		Title:            title,
		CopyNum:          num,
		Author:           author,
		PublicationDate:  publicDate,
		PublisherName:    publisher,
		PublisherAddress: "not provided",
	})

	for i := 1; i <= num; i++ {
		Database.AddCopy(Database.Copy{
			BookID:          id,
			CopyID:          id + i,
			LibraryName:     libName,
			LibraryLocation: libLocation,
		})

	}

	c.IndentedJSON(http.StatusOK, AddBookRet{Success: true})
	return

}
