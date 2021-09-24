package utils

import "github.com/golang-jwt/jwt"

type JwtCustomClaimsUser struct {
	ID    string `json:"id"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}
