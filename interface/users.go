package Interface

import (
	"bookms/Database"
	"strconv"
	"strings"

	"net/http"

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
	Data    []UserData `json:"data"`
}

func filter(userdata UserData , filterdata UserData) bool{
	if filterdata.UserName != "" {
		if strings.Contains(filterdata.UserName,userdata.UserName) || strings.Contains(userdata.UserName,filterdata.UserName) {
			return true
		}
	}
	if filterdata.UserAddress != "" {
		if strings.Contains(filterdata.UserAddress,userdata.UserAddress) || strings.Contains(userdata.UserAddress,filterdata.UserAddress) {
			return true
		}
	}
	if filterdata.UserPhone != "" {
		if strings.Contains(filterdata.UserPhone,userdata.UserPhone) || strings.Contains(userdata.UserPhone,filterdata.UserPhone) {
			return true
		}
	}
	if filterdata.UserName == "" && filterdata.UserAddress == "" && filterdata.UserPhone == ""{
		return true
	}
	return false
}

func GetUsers(c *gin.Context) {

	curPage,_ := strconv.Atoi(c.Query("curPage"))
	userName := c.Query("userName")
	userAddress := c.Query("userAddress")
	userPhone := c.Query("userPhone")
	filterData := UserData{UserName: userName, UserAddress: userAddress, UserPhone: userPhone}

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
				userData := UserData{UserName: value[0], Password: value[1], UserAddress: value[3], UserPhone: value[4]}
				if filter(userData,filterData) {
					data = append(data, userData)
				}
			}
		}

		result := []UserData{}
		if curPage*20 > len(data){
			result = data[(curPage-1)*20:len(data)]
		} else{
			result = data[(curPage-1)*20:curPage*20]
		}
		c.IndentedJSON(http.StatusOK, UsersRet{Success: true, Data: result})
		return
	} else if property == Database.Faculty {
		rows, err := Database.GetAllUsers()
		if err != nil {
			c.IndentedJSON(http.StatusOK, ErrorRes{Success: "false", Msg: err.Error()})
			return
		}
			data := []UserData{}

			for _, value := range rows {

				userData := UserData{UserName: value[0], Password: value[1], UserAddress: value[3], UserPhone: value[4]}
				if filter(userData,filterData) {
					data = append(data, userData)
				}
		}
		result := []UserData{}
		if curPage*20 > len(data){
			result = data[(curPage-1)*20:len(data)]
		} else{
			result = data[(curPage-1)*20:curPage*20]
		}
		c.IndentedJSON(http.StatusOK, UsersRet{Success: true, Data: result})
		return
	} else {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: "false", Msg: "error: no Property found"})
		return
	}

}