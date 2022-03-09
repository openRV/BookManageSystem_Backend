package Interface

import (
	Database "bookms/Database"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

// some

// HandleUploadFile 上传单个文件
func HandleUploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件上传失败"})
		return
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
		return
	}

	fmt.Println(header.Filename)
	fmt.Println(string(content))
	c.JSON(http.StatusOK, gin.H{"msg": "上传成功"})
}

// HandleUploadMutiFile 上传多个文件
func HandleUploadMutiFile(c *gin.Context) {

	// 限制上传文件大小
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 4<<20)

	// 限制放入内存的文件大小
	err := c.Request.ParseMultipartForm(4 << 20)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
		return
	}
	formdata := c.Request.MultipartForm
	files := formdata.File["file"]

	for _, v := range files {
		file, err := v.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
			return
		}
		defer file.Close()

		content, err := ioutil.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
			return
		}

		fmt.Println(v.Filename)
		fmt.Println(string(content))
	}

	c.JSON(http.StatusOK, gin.H{"msg": "上传成功"})
}

// HandleDownloadFile 下载文件

func HandleDownloadFile(c *gin.Context) {

	paperid := c.Query("paperId")
	fmt.Println("PaperId: " + paperid)

	claim, err := VertifyToken(c)
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}

	authName := claim.Name
	authPass := claim.Password

	result, err := Database.SearchPaper(paperid, "")
	if err != nil {
		fmt.Println(err)
		return
	}

	info := Database.DownloadInfo{
		Username:     authName,
		Userpassword: authPass,
		Paperid:      paperid,
		Papertitle:   result[0][2],
		Paperauthor:  result[0][1],
		Downloaddate: time.Now().Format("2006-01-02 15:04:05")}
	err = Database.InsertDownload(info)
	if err != nil {
		fmt.Println(err)
		return
	}

	content := []byte(result[0][6])

	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", "attachment; filename="+result[0][2]+".pdf")
	c.Header("Content-Type", "application/pdf")
	c.Header("Accept-Length", fmt.Sprintf("%d", len(content)))
	c.Writer.Write([]byte(content))
}
