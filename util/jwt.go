package util

import (
	"github.com/golang-jwt/jwt"
)

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

func SignWithClaims(secret []byte) (string, error) {
	// Create the Claims
	claims := MyCustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
		Foo: "ahaaha",
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
}
