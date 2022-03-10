package Interface

import (
	Database "bookms/Database"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

type AddPaperRet struct {
	Success bool `json:"success"`
}

func AddPaper(c *gin.Context) {
	//文件读取
	file, err := c.FormFile("file")
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}
	//文件保存
	file_path := "D:\\ForCourse\\DatabaseSystem\\BookManageSystem_Backend\\static\\" + file.Filename

	/*ip, err := getClientIp()
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}*/
	//ip := "10.128.132.11:8080"

	//Link := "http://" + ip + "/static/" + file.Filename
	c.SaveUploadedFile(file, file_path)

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
		JournalId1 := fmt.Sprintf("%x", sha256.Sum256([]byte(JournalTitle+JournalEditor+scope)))
		JournalId := JournalId1[0:3]
		PaperId1 := fmt.Sprintf("%x", sha256.Sum256([]byte(PaperAuthor+PaperTitle+publishDate+VolumnNum+VolumnEditor)))
		PaperId := PaperId1[0:3]

		var test2 bool
		test2, _ = Database.SearchPaper3(JournalId)
		if !test2 {
			err = Database.InsertJournal(Database.JournalData{
				JournalId,
				JournalTitle,
				JournalEditor,
				scope})
			if err != nil {
				c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
				fmt.Println(err)
				return
			}
			//某期刊是否空
			var test3 bool
			test3, _ = Database.SearchPaper2(JournalId, VolumnNum)
			if !test3 {
				err = Database.InsertVolumn(Database.VolumnData{
					JournalId,
					VolumnNum,
					VolumnEditor,
					publishDate})
				if err != nil {
					c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
					fmt.Println(err)
					return
				}
			}
		}

		data := Database.PaperData{
			PaperId,
			PaperAuthor,
			PaperTitle,
			JournalId,
			VolumnNum,
			" ",
			file_path,
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
		publishDate := c.PostForm("publicDate")
		publishAddress := c.PostForm("publicAddress")
		ConferenceTitle := c.PostForm("ConferenceTitle")
		ProceedingEditor := c.PostForm("ProceedingEditor")
		isOpen := c.PostForm("isOpen")
		ConferenceId1 := fmt.Sprintf("%x", sha256.Sum256([]byte(ConferenceTitle+ProceedingEditor+publishDate+publishAddress)))
		ConferenceId := ConferenceId1[0:3]
		PaperId1 := fmt.Sprintf("%x", sha256.Sum256([]byte(PaperAuthor+PaperTitle+publishDate)))
		PaperId := PaperId1[0:3]

		//检查相关论文是否存在
		var test1 bool
		test1, err = Database.SearchPaper1(ConferenceId)
		if !test1 {
			err = Database.InsertConference(Database.ConferenceData{
				ConferenceId,
				ConferenceTitle,
				ProceedingEditor,
				publishDate,
				publishAddress})
			if err != nil {
				c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
				fmt.Println(err)
				return
			}

		}

		data := Database.PaperData{
			PaperId,
			PaperAuthor,
			PaperTitle,
			" ",
			" ",
			ConferenceId,
			file_path,
			isOpen}
		err := Database.InsertPaper(data)
		if err != nil {
			c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Reuuuuuuuuuuurn")
	c.IndentedJSON(http.StatusOK, AddPaperRet{Success: true})
}

func getClientIp() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", errors.New("Can not find the client ip address!")
}
