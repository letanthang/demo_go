package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	ClientID int `json:"cid"`
	GroupID  int `json:"gid"`
	jwt.StandardClaims
}

func main() {

	claims := MyClaims{
		ClientID: 123,
		GroupID:  456,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
		},
	}

	token, _ := GetJWTToken(claims)
	fmt.Println(token)

	fmt.Println(ParseJWTToken(token))
}

func GetJWTToken(c MyClaims) (string, error) {
	mySigningKey := []byte("MySecret1234")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	ss, err := token.SignedString(mySigningKey)
	return ss, err
}

func ParseJWTToken(tokenString string) *MyClaims {
	mySigningKey := []byte("MySecret1234")

	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(*jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return nil
	}

	if claims, ok := token.Claims.(*MyClaims); ok {
		return claims
	}

	return nil
}
