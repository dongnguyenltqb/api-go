package util

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

func SignWithClaims(claims jwt.Claims, secret []byte) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
}

func JwtParse(destination *jwt.MapClaims, token string, secret []byte) error {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	if !jwtToken.Valid {
		return errors.New("invalid token")
	}
	*destination = jwtToken.Claims.(jwt.MapClaims)
	return nil
}
