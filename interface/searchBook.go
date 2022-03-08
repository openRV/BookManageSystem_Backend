package Interface

import (
	"bookms/Database"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type Copy struct {
	CopyID      string `json:"copyId"`
	LibName     string `json:"libName"`
	LibLocation string `json:"libLocation"`
	Checkout    bool   `json:"checkout"`
}

type BookData struct {
	Bookid     string `json:"bookId"`
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
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}

	authName := claim.Name
	authPass := claim.Password

	_, err = Database.GetUserProperty(Database.User{Username: authName, Password: authPass})
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}

	allBooks, err := Database.GetAllBook()
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}

	data := []BookData{}
	for _, value := range allBooks {
		id, _ := strconv.Atoi(value[0])
		num, _ := strconv.Atoi(value[2])
		bookData := BookData{
			Bookid:     value[0],
			Title:      value[1],
			Num:        num,
			Author:     value[3],
			Publisher:  value[5],
			PublicDate: value[4],
		}
		if bookfilter(bookData, filterTemplate) {
			copies, err := Database.GetAllCopy(Database.Book{ID: id, Title: bookData.Title})
			if err != nil {
				c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
				fmt.Println(err)
				return
			}
			for _, value := range copies {
				copyID := value[1]
				copyIDInt, _ := strconv.Atoi(value[1])
				isBorrowed := false
				borrowed, err := Database.GetBorrowed(Database.Book{ID: id})
				if err != nil {
					c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
					fmt.Println(err)
					return
				}
				for _, value := range borrowed {
					if value.CopyID == copyIDInt {
						isBorrowed = true
					}
				}
				copy := Copy{CopyID: copyID, LibName: value[2], LibLocation: value[3], Checkout: !isBorrowed}
				bookData.Copy = append(bookData.Copy, copy)
			}
			data = append(data, bookData)
		}
	}

	ret := BooksRet{Success: true, Total: len(data)}

	if curPage*10 > len(data) {
		ret.Data = data[(curPage-1)*10:]
	} else {
		ret.Data = data[(curPage-1)*10 : curPage*10]
	}
	c.IndentedJSON(http.StatusOK, ret)

	return
}
