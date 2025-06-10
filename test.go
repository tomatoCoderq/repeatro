package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
    UserID string `json:"user_id"`
    jwt.RegisteredClaims // includes exp, nbf, iat, etc.
}

func main() {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println("here we go error")
	}
	claims := CustomClaims{
		UserID: "12345",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodES256,
		claims)
	token, err := t.SignedString(key)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(token)

	claimsGet := &CustomClaims{}
	parser := jwt.NewParser(jwt.WithLeeway(5*time.Minute))
	tokenGot, err := parser.ParseWithClaims(token, claimsGet, func(token *jwt.Token) (interface{}, error) {
        return &key.PublicKey, nil
    })
	if err != nil {
		fmt.Println("err")
	}
	if !tokenGot.Valid {
		fmt.Println("invalid token")
	}
	fmt.Println(claimsGet.UserID)

}
