package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("blablaest")

type Claim struct {
	UserId string
	Role   string
	jwt.RegisteredClaims
}

func GenerateToken(userId string, role string) (string, error) {
	now := time.Now()

	claims := Claim{
		UserId: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userId,
			Issuer:    "bluedit-api",
			Audience:  []string{"bluedit-front"},
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(4 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
