package Interface

import (
	"bookms/Database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type Copy struct {
	CopyID      int    `json:"copyID"`
	LibName     string `json:"libName"`
	LibLocation string `json:"libLocation"`
	Checkout    bool   `json:"checkout"`
}

type BookData struct {
	Bookid     int    `json:"bookId"`
	Title      string `json:"title"`
	Num        int    `json:"num"`
	Author     string `json:"author"`
	Publisher  string `json:"publisher"`
	PublicDate string `json:"publicDate"`
	Copy       []Copy `json:"copy"`
}

type BooksRet struct {
	Success bool       `json:"success"`
	Data    []BookData `json:"data"`
	Total   int        `json:"total"`
}

func bookfilter(searchBook BookData, filterBook BookData) bool {
	if filterBook.Title != "" {
		if (strings.Contains(searchBook.Title, filterBook.Title)) || (strings.Contains(filterBook.Title, searchBook.Title)) {
			return true
		}
	}
	if filterBook.Author != "" {
		if (strings.Contains(searchBook.Author, filterBook.Author)) || (strings.Contains(filterBook.Author, searchBook.Author)) {
			return true
		}
	}
	if filterBook.Publisher != "" {
		if (strings.Contains(searchBook.Publisher, filterBook.Publisher)) || (strings.Contains(filterBook.Publisher, searchBook.Publisher)) {
			return true
		}
	}

	if (filterBook.Title == "") && (filterBook.Author == "") && (filterBook.Publisher == "") {
		return true
	}
	return false

}

func SearchBook(c *gin.Context) {
	curPage, _ := strconv.Atoi(c.Query("curPage"))
	title := c.Query("title")
	author := c.Query("author")
	publisher := c.Query("publisher")

	filterTemplate := BookData{Title: title, Author: author, Publisher: publisher}

	if curPage < 1 {
		curPage = 1
	}
	claim, err := VertifyToken(c)
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: "false", Msg: err.Error()})
		return
	}

	authName := claim.Name
	authPass := claim.Password

	_, err = Database.GetUserProperty(Database.User{Username: authName, Password: authPass})
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: "false", Msg: err.Error()})
		return
	}

	allBooks, err := Database.GetAllBook()
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: "false", Msg: err.Error()})
		return
	}

	data := []BookData{}
	for _, value := range allBooks {
		id, _ := strconv.Atoi(value[0])
		num, _ := strconv.Atoi(value[2])
		bookData := BookData{
			Bookid:     id,
			Title:      value[1],
			Num:        num,
			Author:     value[3],
			Publisher:  value[5],
			PublicDate: value[4],
		}
		if bookfilter(bookData, filterTemplate) {
			copies, err := Database.GetAllCopy(Database.Book{ID: bookData.Bookid})
			if err != nil {
				c.IndentedJSON(http.StatusOK, ErrorRes{Success: "false", Msg: err.Error()})
				return
			}
			for _, value := range copies {
				copyID, _ := strconv.Atoi(value[1])
				// TODO search database for copy's checkout
				copy := Copy{CopyID: copyID, LibName: value[2], LibLocation: value[3], Checkout: true}
				bookData.Copy = append(bookData.Copy, copy)
			}
			data = append(data, bookData)
		}
	}

	ret := BooksRet{Success: true, Total: (len(data) / 20) + 1}

	if curPage*20 > len(data) {
		ret.Data = data[(curPage-1)*20:]
	} else {
		ret.Data = data[(curPage-1)*20 : curPage*20]
	}
	c.IndentedJSON(http.StatusOK, ret)

	return
}
