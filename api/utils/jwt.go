package utils

import "github.com/golang-jwt/jwt"

type JwtCustomClaimsUser struct {
	ID    uint64 `json:"id"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}
