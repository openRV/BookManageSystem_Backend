package Interface

import (
	"bookms/Database"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DelPaperRet struct {
	Success bool `json:"success"`
}

func DeletePaper(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)

	paperid := json["paperId"].(string)

	fmt.Println("PapperId From Web: " + paperid)
	claim, err := VertifyToken(c)
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}

	authName := claim.Name
	authPass := claim.Password

	property, err := Database.GetUserProperty(Database.User{Username: authName, Password: authPass})
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}

	if property == Database.Student {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: "your property is student, not able to delete book"})
		return
	}

	err = Database.DeletePaper(paperid)
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: false, Msg: err.Error()})
		fmt.Println(err)
		return
	}
	c.IndentedJSON(http.StatusOK, DelPaperRet{Success: true})
}
