package modelutil

import "github.com/golang-jwt/jwt/v5"

type JwtTokenClaims struct {
	jwt.RegisteredClaims

	// apa lagi yang akan dibuat di payload
	UserId   string   `json:"userId"`
	Role     string   `json:"role"`
	Services []string `json:"services"`
}
