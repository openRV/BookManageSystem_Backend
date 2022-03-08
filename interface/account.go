package Interface

import (
	"bookms/Database"
	"github.com/gin-gonic/gin"
	"net/http"
)

type accountData struct {
	UserName    string `json:"userName"`
	Password    string `json:"password"`
	UserAddress string `json:"userAddress"`
	UserPhone   string `json:"userPhone"`
}

type AccountRet struct {
	Success bool        `json:"success"`
	Data    accountData `json:"data"`
}

func GetAccount(c *gin.Context) {

	claim, err := VertifyToken(c)
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		return
	}

	username := claim.Name
	password := claim.Password

	User, err := Database.SearchUser(Database.User{Username: username, Password: password})
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		return
	}
	data := accountData{UserName: User.Username, Password: User.Password, UserAddress: User.Address, UserPhone: User.Phone}
	c.IndentedJSON(http.StatusOK, AccountRet{Success: true, Data: data})
}

func ModifyAccount(c *gin.Context) {
	claim, err := VertifyToken(c)
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		return
	}

	json := make(map[string]string)
	c.BindJSON(&json)

	userName := json["userName"]
	userAddress := json["userAddress"]
	userPhone := json["userPhone"]
	changePassword := json["changePassword"]
	oldPassword := json["oldPassword"]
	newPassword := json["newPassword"]

	username := claim.Name
	password := claim.Password

	if changePassword == "true" && oldPassword != password {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: "old password dismatch"})
		return
	}
	err = Database.DeleteUser(Database.User{Username: username, Password: password})

	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		return
	}

	finalPassword := password
	if changePassword == "true" {
		finalPassword = newPassword
	}

	err = Database.RegisterUser(Database.User{Username: userName, Address: userAddress, Phone: userPhone, Password: finalPassword, Property: Database.Student})
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, AccountRet{Success: true, Data: accountData{UserName: userName, UserAddress: userAddress, UserPhone: userPhone, Password: finalPassword}})
}
