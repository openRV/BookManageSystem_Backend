package Interface

import (
	"github.com/gin-gonic/gin"
)

type AddBookRet struct {
	Success bool `json:"success"`
}

type BookCopy struct {
	libName     string `json:"libName"`
	libLocation string `json:"libLocation"`
	checkout    bool   `json:"checkout"`
}

func AddBook(c *gin.Context) {
	//title := c.PostForm("title")
	//author := c.PostForm("author")
	//num := c.PostForm("num")
	//publisher := c.PostForm("publisher")
	//publicDate := c.PostForm("publicDate")
	//copy := c.PostFormArray("copy")

	//var data []BookCopy
	//json.Unmarshal([]byte(copy), &data)
}
