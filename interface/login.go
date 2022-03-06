package Interface

import (
	"net/http"

	"bookms/Database"

	"github.com/gin-gonic/gin"
)

type userRet struct {
	Success string `json:"success"`
	Token   string `json:"token"`
}

func PostUser(c *gin.Context) {

	username := c.PostForm("userName")
	password := c.PostForm("password")

	token := GetToken(c)

	_, err := Database.SearchUser(Database.User{Username: username, Password: password})
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: "false", Msg: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, userRet{Success: "true", Token: token})

}
