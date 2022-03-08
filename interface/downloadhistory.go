package Interface

import (
	"bookms/Database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DownloadHistoryInfo struct {
	PaperId      string `json:"paperId"`
	Author       string `json:"author"`
	PaperTitle   string `json:"title"`
	DownloadDate string `json:"downloadDate"`
}

type DownloadHistoryRet struct {
	Success bool                  `json:"success"`
	Data    []DownloadHistoryInfo `json:"data"`
	Total   int
}

func SearchDownH(c *gin.Context) {
	curPage, _ := strconv.Atoi(c.Query("curPage"))
	if curPage < 1 {
		curPage = 1
	}

	claim, err := VertifyToken(c)
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		return
	}

	authName := claim.Name
	authPass := claim.Password

	history, err := Database.SearchDownload(authName, authPass)
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		return
	}

	var data []DownloadHistoryInfo
	for _, value := range history {
		data = append(data, DownloadHistoryInfo{
			PaperId:      value[0],
			Author:       value[2],
			PaperTitle:   value[1],
			DownloadDate: value[3],
		})
	}

	ret := DownloadHistoryRet{
		Success: true,
		Total:   len(data)}

	if curPage*10 > len(data) {
		ret.Data = data[(curPage-1)*10:]
	} else {
		ret.Data = data[(curPage-1)*10 : curPage*10]
	}

	c.IndentedJSON(http.StatusOK, ret)
	return
}
