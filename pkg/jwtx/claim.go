package jwtx

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Email  string `json:"username"`
	UserID int    `json:"user_id"`
	jwt.StandardClaims
}
