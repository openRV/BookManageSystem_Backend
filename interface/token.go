package Interface

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secret = []byte("bookmanagesystem")

var effectTime = 24 * time.Hour

type Claim struct {
	Name     string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GetToken(json map[string]string) string {

	expireTime := time.Now().Add(effectTime)
	claims := &Claim{
		Name:     json["username"],
		Password: json["password"],
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "localhost",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

func parseToken(tokenString string) (*jwt.Token, *Claim, error) {
	Claim := &Claim{}
	token, err := jwt.ParseWithClaims(tokenString, Claim, func(tokenString *jwt.Token) (i interface{}, err error) {
		return secret, nil
	})
	return token, Claim, err
}

func VertifyToken(ctx *gin.Context) (*Claim, error) {

	tokenString := ctx.GetHeader("Authorization")

	if tokenString == "" {
		return nil, errors.New("empty tokenString")
	}
	token, claim, err := parseToken(tokenString)
	if err != nil || !token.Valid {
		return nil, errors.New("Invalid Token")
	}

	fmt.Println("Parsed token: name:", claim.Name, " password:", claim.Password)
	return claim, nil

}
