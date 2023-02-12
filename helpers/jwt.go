package helpers

import "github.com/golang-jwt/jwt/v4"

var JWTKey = []byte("jqlkjdhuahfjkalhdja1jk12i31231")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
