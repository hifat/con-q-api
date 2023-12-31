package token

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/hifat/con-q-api/internal/app/domain/authDomain"
)

type AuthClaims struct {
	Passport authDomain.Passport
	jwt.RegisteredClaims
}
