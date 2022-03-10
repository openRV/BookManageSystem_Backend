package Interface

import (
	"bookms/Database"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//Conference
type ConferencePaperData struct {
	PaperId          string `json:"paperId"`
	PaperTitle       string `json:"title"`
	PaperAuthor      string `json:"author"`
	PublishDate      string `json:"publicDate"`
	PublishAddress   string `json:"publicAddress"`
	ConferenceTitle  string `json:"ConferenceTitle"`
	ProceedingEditor string `json:"ProceedingEditor"`
}

type ConferencePaperRet struct {
	Success bool                  `json:"success"`
	Data    []ConferencePaperData `json:"data"`
	Total   int                   `json:"total"`
}

func SearchConferencePaper(c *gin.Context) {
	curPage, _ := strconv.Atoi(c.Query("curPage"))
	title := c.Query("title")
	author := c.Query("author")
	ConferenceTitle := c.Query("ConferenceTitle")
	ProceedingEditor := c.Query("ProceedingEditor")

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

	result, err := Database.SearchConferencePaper(title, author, ConferenceTitle, ProceedingEditor)
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}

	data := []ConferencePaperData{}
	for _, value := range result {
		paperData := ConferencePaperData{
			PaperId:          value[0],
			PaperTitle:       value[1],
			PaperAuthor:      value[2],
			PublishDate:      value[3],
			PublishAddress:   value[4],
			ConferenceTitle:  value[5],
			ProceedingEditor: value[6],
		}
		data = append(data, paperData)
	}

	ret := ConferencePaperRet{
		Success: true,
		Total:   len(data)}

	if curPage*10 > len(data) {
		ret.Data = data[(curPage-1)*10:]
	} else {
		ret.Data = data[(curPage-1)*10 : curPage*10]
	}
	c.IndentedJSON(http.StatusOK, ret)
}

//Journal
type JournalPaperDate struct {
	PaperId       string `json:"paperId"`
	PaperTitle    string `json:"title"`
	PaperAuthor   string `json:"author"`
	PublishDate   string `json:"publicationDate"`
	Scope         string `json:"scope"`
	JournalTitle  string `json:"JournalTtile"`
	JournalEditor string `json:"JournalEditor"`
	VolumnNum     string `json:"VolumeNum"`
	VolumnEditor  string `json:"VolumeEditor"`
}

type JournalPaperRet struct {
	Success bool               `json:"success"`
	Data    []JournalPaperDate `json:"data"`
	Total   int
}

func SearchJournalPaper(c *gin.Context) {
	curPage, _ := strconv.Atoi(c.Query("curPage"))
	title := c.Query("title")
	scope := c.Query("scope")
	author := c.Query("author")
	JournalTitle := c.Query("JournalTitle")
	VolumeNum := c.Query("VolumeNum")
	VolumeEditor := c.Query("VolumeEditor")

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

	result, err := Database.SearchJournalPaper(title, author, JournalTitle, scope, VolumeNum, VolumeEditor)
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}

	data := []JournalPaperDate{}
	for _, value := range result {
		paperData := JournalPaperDate{
			PaperId:       value[0],
			PaperTitle:    value[1],
			PaperAuthor:   value[2],
			PublishDate:   value[3],
			Scope:         value[4],
			JournalTitle:  value[5],
			JournalEditor: value[6],
			VolumnNum:     value[7],
			VolumnEditor:  value[8],
		}
		data = append(data, paperData)
	}

	ret := JournalPaperRet{
		Success: true,
		Total:   len(data)}

	if curPage*10 > len(data) {
		ret.Data = data[(curPage-1)*10:]
	} else {
		ret.Data = data[(curPage-1)*10 : curPage*10]
	}
	c.IndentedJSON(http.StatusOK, ret)
}

//OpenPaper

type OpenPaperInfo struct {
	PaperId     string `json:"paperId"`
	PaperTitle  string `json:"title"`
	PaperAuthor string `json:"author"`
	PublishDate string `json:"publicDate"`
	Link        string `json:"link"`
}

type OpenPaperRet struct {
	Success bool            `json:"success"`
	Data    []OpenPaperInfo `json:"data"`
	Total   int
}

func GetOpenPaper(c *gin.Context) {
	curPage, _ := strconv.Atoi(c.Query("curPage"))
	fmt.Println("CurPage:", curPage)
	if curPage < 1 {
		curPage = 1
	}
	/*
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
		}*/

	result, err := Database.GetOpenPaper()
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}

	data := []OpenPaperInfo{}
	for _, value := range result {
		file_slice := strings.Split(value[4], "\\")
		file_name := file_slice[len(file_slice)-1]
		ip := "10.128.132.11"
		downloadLink := "http://" + ip + ":8080/static" + file_name
		paperData := OpenPaperInfo{
			PaperId:     value[0],
			PaperTitle:  value[1],
			PaperAuthor: value[2],
			PublishDate: value[3],
			Link:        downloadLink,
		}
		data = append(data, paperData)
	}

	ret := OpenPaperRet{
		Success: true,
		Total:   len(data)}

	if curPage*10 > len(data) {
		ret.Data = data[(curPage-1)*10:]
	} else {
		ret.Data = data[(curPage-1)*10 : curPage*10]
	}
	c.IndentedJSON(http.StatusOK, ret)
}
