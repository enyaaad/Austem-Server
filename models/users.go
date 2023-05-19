package models

import "github.com/dgrijalva/jwt-go/v4"

type Users struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Password string `json:"password"`
}
