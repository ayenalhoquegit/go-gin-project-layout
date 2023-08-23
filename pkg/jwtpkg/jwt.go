package jwtpkg

import (
	"errors"
	"time"

	"github.com/ayenalhoquegit/go-gin-project-layout/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)

// Create the JWT key used to create the signature
var JwtKey = []byte(config.GetEnvValue("JWT_SECRET"))

// jwt.RegisteredClaims is an embedded type
type Payload struct {
	Id int `json:"id"`
}

// jwt.RegisteredClaims is an embedded type
type Claims struct {
	Payload Payload `json:"payload"`
	jwt.RegisteredClaims
}

func GenerateToken(payload Payload) string {
	/* RegisteredClaims: jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(datetime.Now()),
		NotBefore: jwt.NewNumericDate(datetime.Now()),
		Issuer:    "test",
		Subject:   "somebody",
		ID:        "1",
		Audience:  []string{"somebody_else"},
	}, */
	// Declare the expiration time of the token
	expirationTime := time.Now().AddDate(0, 0, 1)
	claims := &Claims{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(JwtKey)
	return tokenString
}

func VerifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		},
	)
	if err != nil {
		return nil, errors.New("malformed token")
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil

}
