package Interface

import (
	"net/http"
	"strconv"
	"fmt"
	"bookms/Database"

	"github.com/gin-gonic/gin"
)

type UserData struct {
	UserID      string `json:"userID"`
	UserName    string `json:"userName"`
	Password    string `json:"password"`
	UserAddress string `json:"userAddress"`
	UserPhone   string `json:"userPhone"`
}

type UsersRet struct {
	Success bool `json:"success"`
	Data    []UserData
}

func GetUsers(c *gin.Context) {

	curPage,_ := strconv.Atoi(c.Query("curPage"))
	userName := c.Query("userName")
	userAddress := c.QUery("userAddress")
	userPhone := c.Query("userPhone")

	// TODO filter result using following values

	fmt.Println(curPage)

	claim, err := VertifyToken(c)
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: "false", Msg: err.Error()})
		return
	}

	authName := claim.Name
	authPass := claim.Password

	property, err := Database.GetUserProperty(Database.User{Username: authName, Password: authPass})
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: "false", Msg: err.Error()})
		return
	}

	if property == Database.Student {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: "false", Msg: err.Error()})
		return
	} else if property == Database.Staff {
		rows, err := Database.GetAllUsers()
		if err != nil {
			c.IndentedJSON(http.StatusOK, ErrorRes{Success: "false", Msg: err.Error()})
			return
		}

		data := []UserData{}

		for _, value := range rows {
			if value[2] == strconv.Itoa(Database.Student) {
				data = append(data, UserData{UserName: value[0], Password: value[1], UserAddress: value[3], UserPhone: value[4]})
			}
		}
		c.IndentedJSON(http.StatusOK, UsersRet{Success: true, Data: data})
		return
	} else if property == Database.Faculty {
		rows, err := Database.GetAllUsers()
		if err != nil {
			c.IndentedJSON(http.StatusOK, ErrorRes{Success: "false", Msg: err.Error()})
			return
		}
		data := []UserData{}

		for _, value := range rows {
				data = append(data, UserData{UserName: value[0], Password: value[1], UserAddress: value[3], UserPhone: value[4]})
		}
		c.IndentedJSON(http.StatusOK, UsersRet{Success: true, Data: data})
		return
	} else {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: "false", Msg: "error: no Property found"})
		return
	}

}
