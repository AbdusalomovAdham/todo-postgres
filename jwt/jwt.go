package jwt

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret_key")

type Claims struct {
	Uid      string
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func SignToken(uid, username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Hour)
	cliams := &Claims{
		Uid:      uid,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "my-app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(authorization string) (*Claims, error) {
	token := strings.TrimPrefix(authorization, "Bearer ")
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		return jwtKey, nil
	})

	if err != nil || !tkn.Valid {
		return nil, err
	}
	return claims, nil
}
