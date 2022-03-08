package Interface

import (
	"bookms/Database"
	"fmt"
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
		fmt.Println(err)
		return
	}

	username := claim.Name
	password := claim.Password

	User, err := Database.SearchUser(Database.User{Username: username, Password: password})
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}
	data := accountData{UserName: User.Username, Password: User.Password, UserAddress: User.Address, UserPhone: User.Phone}
	c.IndentedJSON(http.StatusOK, AccountRet{Success: true, Data: data})
}

func ModifyAccount(c *gin.Context) {
	claim, err := VertifyToken(c)
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}

	json := make(map[string]interface{})
	c.BindJSON(&json)

	userName := json["userName"].(string)
	userAddress := json["userAddress"].(string)
	userPhone := json["userPhone"].(string)
	changePassword := json["changePassword"].(bool)
	oldPassword := json["oldPassword"].(string)
	newPassword := json["newPassword"].(string)

	username := claim.Name
	password := claim.Password

	if changePassword && oldPassword != password {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: "old password dismatch"})
		fmt.Println(err)
		return
	}
	err = Database.DeleteUser(Database.User{Username: username, Password: password})

	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}

	finalPassword := password
	if changePassword {
		finalPassword = newPassword
	}

	err = Database.RegisterUser(Database.User{Username: userName, Address: userAddress, Phone: userPhone, Password: finalPassword, Property: Database.Student})
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}
	c.IndentedJSON(http.StatusOK, AccountRet{Success: true, Data: accountData{UserName: userName, UserAddress: userAddress, UserPhone: userPhone, Password: finalPassword}})
}
