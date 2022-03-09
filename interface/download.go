package Interface

import (
	"bookms/Database"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
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
	result, err := Database.SearchPaper(paperid, "")
	if err != nil {
		fmt.Println(err)
		return
	}

	content := "hello world, 我是一个文件，" + result[0][6]

	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", "attachment; filename=hello.pdf")
	c.Header("Content-Type", "application/text/plain")
	c.Header("Accept-Length", fmt.Sprintf("%d", len(content)))
	c.Writer.Write([]byte(content))
}
