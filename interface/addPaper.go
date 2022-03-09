package Interface

import (
	Database "bookms/Database"
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddPaperRet struct {
	Success bool `json:success`
}

func AddPaper(c *gin.Context) {
	//文件读取
	read, err := c.FormFile("file")
	reader, _ := read.Open()
	file := make([]byte, 10)
	_, err = reader.Read(file)
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}

	//权限认证
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
	//分类处理Journal和Paper
	isJournal := c.PostForm("isJournal")
	if isJournal == "true" {
		PaperTitle := c.PostForm("title")
		PaperAuthor := c.PostForm("author")
		publishDate := c.PostForm("publicationDate")
		scope := c.PostForm("scope")
		JournalTitle := c.PostForm("JournalTitle")
		JournalEditor := c.PostForm("JournalEditor")
		VolumnNum := c.PostForm("VolumeNum")
		VolumnEditor := c.PostForm("VolumeEditor")
		isOpen := c.PostForm("isOpen")
		JournalId := fmt.Sprintf("%x", sha256.Sum256([]byte(JournalTitle+JournalEditor+scope)))
		PaperId := fmt.Sprintf("%x", sha256.Sum256([]byte(PaperAuthor+PaperTitle+publishDate+VolumnNum+VolumnEditor)))

		var test2 bool
		test2, err = Database.SearchPaper2(JournalId, VolumnNum)
		if !test2 {
			Database.InsertVolumn(Database.VolumnData{
				JournalId,
				VolumnNum,
				VolumnEditor,
				publishDate})
			//某期刊是否空
			var test3 bool
			test3, err = Database.SearchPaper3(JournalId)
			if !test3 {
				Database.InsertJournal(Database.JournalData{
					JournalId,
					JournalTitle,
					JournalEditor,
					scope})
			}
		}

		data := Database.PaperData{
			PaperId,
			PaperAuthor,
			PaperTitle,
			JournalId,
			VolumnNum,
			" ",
			string(file),
			isOpen}
		err := Database.InsertPaper(data)
		if err != nil {
			c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
			fmt.Println(err)
			return
		}
	} else {
		PaperTitle := c.PostForm("title")
		PaperAuthor := c.PostForm("author")
		publishDate := c.PostForm("publicationDate")
		publishAddress := c.PostForm("publicAddress")
		ConferenceTitle := c.PostForm("ConferenceTitle")
		ProceedingEditor := c.PostForm("ProceedingEditor")
		isOpen := c.PostForm("isOpen")
		ConferenceId := fmt.Sprintf("%x", sha256.Sum256([]byte(ConferenceTitle+ProceedingEditor+publishDate+publishAddress)))
		PaperId := fmt.Sprintf("%x", sha256.Sum256([]byte(PaperAuthor+PaperTitle+publishDate)))

		//检查相关论文是否存在
		var test1 bool
		test1, err = Database.SearchPaper1(ConferenceId)
		if !test1 {
			Database.InsertConference(Database.ConferenceData{
				ConferenceId,
				ConferenceTitle,
				ProceedingEditor,
				publishDate,
				publishAddress})
		}

		data := Database.PaperData{
			PaperId,
			PaperAuthor,
			PaperTitle,
			" ",
			" ",
			ConferenceId,
			string(file),
			isOpen}
		err := Database.InsertPaper(data)
		if err != nil {
			c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
			fmt.Println(err)
			return
		}
	}
	c.IndentedJSON(http.StatusOK, AddPaperRet{Success: true})
}
