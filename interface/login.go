package Interface

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userRet struct {
	Success string `json:"success"`
	Token   string `json:"token"`
}

func PostUser(c *gin.Context) {

	// GET parameters
	username := c.Query("userName")
	password := c.Query("password")

	token := GetToken(c)

	fmt.Println("getting  username:", username, " password:", password, "token: ", token)

	c.IndentedJSON(http.StatusOK, userRet{Success: "true", Token: token})

}
