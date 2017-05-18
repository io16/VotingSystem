package config

import "github.com/dgrijalva/jwt-go"

type JwtCustomClaims struct {
	Name string
	Role string
	ID uint
	jwt.StandardClaims
}