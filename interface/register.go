package Interface

import (
	"bookms/Database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRet struct {
	Success bool
}

func RegisterUser(c *gin.Context) {

	json := make(map[string]string)
	c.BindJSON(&json)

	username := json["username"]
	password := json["password"]
	userAddress := json["useraddress"]
	userPhone := json["userphone"]

	fmt.Println(username, password, userAddress, userPhone)

	err := Database.RegisterUser(Database.User{Username: username, Password: password, Address: userAddress, Phone: userPhone, Property: Database.Student})
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, RegisterRet{Success: true})

}
