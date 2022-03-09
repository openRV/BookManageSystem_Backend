package Interface

import (
	"bookms/Database"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BorrowRet struct {
	Success bool `json:"success"`
}

type BorrowInfo struct {
	BookId     int    `json:"bookId"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Publisher  string `json:"publisher"`
	BorrowDate string `json:"borrowDate"`
	ReturnDate string `json:"returnDate"`
}

type BorrowInfoRet struct {
	Success bool         `json:"success"`
	Total   int          `json:"total"`
	Data    []BorrowInfo `json:"data"`
}

func BorrowBook(c *gin.Context) {

	json := make(map[string]interface{})
	c.BindJSON(&json)

	bookId, _ := strconv.Atoi(json["bookId"].(string))

	fmt.Println(bookId)

	claim, err := VertifyToken(c)
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}

	authName := claim.Name
	authPass := claim.Password

	err = Database.Borrow(Database.Book{ID: bookId}, Database.User{Username: authName, Password: authPass})
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}

	c.IndentedJSON(http.StatusOK, BorrowRet{Success: true})
	return
}

func ReturnBook(c *gin.Context) {

	json := make(map[string]interface{})
	c.BindJSON(&json)

	bookId := int(json["bookId"].(float64))

	fmt.Println(bookId)

	claim, err := VertifyToken(c)
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}

	authName := claim.Name
	authPass := claim.Password

	err = Database.Return(Database.Book{ID: bookId}, Database.User{Username: authName, Password: authPass})
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}

	c.IndentedJSON(http.StatusOK, BorrowRet{Success: true})
	return
}

func Borrowing(c *gin.Context) {
	curPage, _ := strconv.Atoi(c.Query("curPage"))
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

	borrowing, err := Database.GetBorrowingBy(Database.User{Username: authName, Password: authPass})
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}
	var data []BorrowInfo
	for _, value := range borrowing {
		data = append(data, BorrowInfo{
			BookId:     value.BookId,
			Title:      value.Title,
			Author:     value.Author,
			Publisher:  value.Publisher,
			BorrowDate: value.BorrowDate,
			ReturnDate: value.ReturnDate,
		})
	}

	var result []BorrowInfo
	if curPage*10 > len(data) {
		result = data[(curPage-1)*10:]
	} else {
		result = data[(curPage-1)*10 : curPage*10]
	}
	c.IndentedJSON(http.StatusOK, BorrowInfoRet{Success: true, Data: result, Total: len(data)})

}

func Borrowed(c *gin.Context) {

	curPage, _ := strconv.Atoi(c.Query("curPage"))
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

	borrowing, err := Database.GetBorrowedBy(Database.User{Username: authName, Password: authPass})
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}
	var data []BorrowInfo
	for _, value := range borrowing {
		data = append(data, BorrowInfo{
			BookId:     value.BookId,
			Title:      value.Title,
			Author:     value.Author,
			Publisher:  value.Publisher,
			BorrowDate: value.BorrowDate,
			ReturnDate: value.ReturnDate,
		})
	}

	var revdata []BorrowInfo

	for i := len(data) - 1; i >= 0; i-- {
		revdata = append(revdata, data[i])
	}

	var result []BorrowInfo
	if curPage*10 > len(revdata) {
		result = revdata[(curPage-1)*10:]
	} else {
		result = revdata[(curPage-1)*10 : curPage*10]
	}
	c.IndentedJSON(http.StatusOK, BorrowInfoRet{Success: true, Data: result, Total: len(revdata)})
}
