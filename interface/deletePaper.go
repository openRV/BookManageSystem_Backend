package Interface

import (
	Database "bookms/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DelPaperRet struct {
	Success bool `json:"success"`
}

func DeletePaper(c *gin.Context) {
	paperid := c.Query("paperId")

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
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: "false", Msg: "your property is student, not able to delete book"})
		return
	}

	err = Database.DeletePaper(paperid)
	if err != nil {
		c.IndentedJSON(http.StatusOK, ErrorRes{Success: "false", Msg: err.Error()})
		return
	}
}
