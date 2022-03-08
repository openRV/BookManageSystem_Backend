package Interface

import (
	"bookms/Database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRet struct {
	Success bool `json:"success"`
}

func RegisterUser(c *gin.Context) {

	json := make(map[string]string)
	c.BindJSON(&json)

	username := json["userName"]
	password := json["password"]
	userAddress := json["userAddress"]
	userPhone := json["userPhone"]

	err := Database.RegisterUser(Database.User{Username: username, Password: password, Address: userAddress, Phone: userPhone, Property: Database.Student})
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}

	c.IndentedJSON(http.StatusOK, RegisterRet{Success: true})

}
