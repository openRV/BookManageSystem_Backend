package Interface

import (
	"net/http"

	"bookms/Database"

	"github.com/gin-gonic/gin"
)

type UserToken struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}

type UserRet struct {
	Success string    `json:"success"`
	Data    UserToken `json:"data"`
}

func PostUser(c *gin.Context) {

	json := make(map[string]string)
	c.BindJSON(&json)

	username := json["username"]
	password := json["password"]

	token := GetToken(json)
	user, err := Database.SearchUser(Database.User{Username: username, Password: password})
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: "false", Msg: err.Error()})
		return
	}

	var str string
	if user.Property == Database.Student {
		str = "STU"
	} else if user.Property == Database.Staff {
		str = "FAC"
	} else {
		str = "ADM"
	}

	c.IndentedJSON(http.StatusOK, UserRet{Success: "true", Data: UserToken{Token: token, Role: str}})

}
