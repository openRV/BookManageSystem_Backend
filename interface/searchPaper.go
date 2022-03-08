package Interface

import (
	"bookms/Database"
	"net/http"
	"strconv"

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
	ProceedingEditor string `json:"proceedingEditor"`
}

type ConferencePaperRet struct {
	Success bool                  `json:"success"`
	Data    []ConferencePaperData `json:"data"`
	Total   int                   `json:"total"`
}

func SearchConferencebook(c *gin.Context) {
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

	result, err := Database.SearchConferencePaper(title, author, ConferenceTitle, ProceedingEditor)
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: "false", Msg: err.Error()})
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

	if curPage*7 > len(data) {
		ret.Data = data[(curPage-1)*7:]
	} else {
		ret.Data = data[(curPage-1)*7 : curPage*7]
	}
	c.IndentedJSON(http.StatusOK, ret)
}

//Journal
type JournalPaperDate struct {
	PaperId       string `json:"paperId"`
	PaperTitle    string `json:"title"`
	PaperAuthor   string `json:"author"`
	PublishDate   string `json:"publicationDate"`
	Scope         string `json:"scopr"`
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

	result, err := Database.SearchJournalPaper(title, author, JournalTitle, scope, VolumeNum, VolumeEditor)
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: "false", Msg: err.Error()})
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

	if curPage*9 > len(data) {
		ret.Data = data[(curPage-1)*9:]
	} else {
		ret.Data = data[(curPage-1)*9 : curPage*9]
	}
	c.IndentedJSON(http.StatusOK, ret)
}
