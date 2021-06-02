package test_jwt

import (
	"fmt"
	"testing"
	"time"

	"github.com/develop1024/jwt-go"
)

func GenerationToken(secret []byte, claims jwt.MapClaims) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(secret)

	return tokenString, err
}

func ParseToken(tokenString string, secret []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func TestJWTGo(t *testing.T) {
	claims := jwt.MapClaims{
		"iss": "wang",
		"exp": time.Now().Add(time.Second * 1).Unix(),
		"foo": "bar",
	}

	token, err := GenerationToken([]byte("abc"), claims)
	if err == nil {
		fmt.Println("token: ", token)
	} else {
		fmt.Println("err: ", err)
	}

	claims, err = ParseToken(token, []byte("ac"))
	if err != nil {
		fmt.Println("parse error: ", err) //
	} else {
		fmt.Println("claims: ", claims)
	}
}
