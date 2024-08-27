package entities

import (
	jwt "github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Role string `json:"role"`
	ID   string `json:"id"`
	jwt.RegisteredClaims
}

// Valid implements jwt.Claims.
func (j *JwtCustomClaims) Valid() error {
	return j.RegisteredClaims.Valid()
}

type JwtCustomRefreshClaims struct {
	ID   string `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

// Valid implements jwt.Claims.
func (j *JwtCustomRefreshClaims) Valid() error {
	return j.RegisteredClaims.Valid()
}
