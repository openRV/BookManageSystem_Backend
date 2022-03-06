package Interface

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Success string `json:"success"`
	Token   string `json:"token"`
}

func GetUser(c *gin.Context) {

	// GET parameters
	username := c.Query("userName")
	password := c.Query("password")

	token := sha256.Sum256([]byte(username + password))
	tokenstr := hex.EncodeToString(token[:])

	fmt.Println("getting  username:", username, " password:", password, "token: ", tokenstr)

	user := User{
		Success: "true",
		Token:   tokenstr,
	}

	c.IndentedJSON(http.StatusOK, user)

}
