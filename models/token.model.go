package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type Token struct {
	UserID uuid.UUID
	jwt.StandardClaims
}
