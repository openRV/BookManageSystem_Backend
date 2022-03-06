package Interface

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

var users = []User{
	{UserName: "a", Password: "123"},
	{UserName: "ab", Password: "123"},
	{UserName: "abc", Password: "123"},
}

func GetUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
