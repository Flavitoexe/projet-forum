package auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(tokenString string) (*Claim, error) {
	claim := &Claim{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claim,
		func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("Méthode de signature inattendue - %v", token.Header["alg"])
			}
			return jwtSecret, nil
		},
		jwt.WithIssuer("bluedit-api"),
		jwt.WithAudience("bluedit-front"),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)

	if err != nil {
		return nil, fmt.Errorf("Erreur forme, expiration ou validité du token - Erreur\n\t%s", err.Error())
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claim, nil
}
