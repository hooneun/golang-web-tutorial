package helpers

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type claims struct {
	ID uint
	jwt.StandardClaims
}

// JWTToken struct
type JWTToken struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int64  `json:"expiresIn"`
	TokenType   string `json:"tokenType"`
}

const jwtExpiredTime = 20 * time.Minute

var jwtSecretKey = []byte("jwtkey")

// GetJWTToken !
func GetJWTToken(id uint) (JWTToken, error) {
	expirationTime := time.Now().Add(jwtExpiredTime).Unix()
	claims := &claims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(jwtSecretKey)

	return JWTToken{
		AccessToken: accessToken,
		ExpiresIn:   expirationTime,
		TokenType:   "bearer",
	}, err
}

// CheckJWTToken !
func CheckJWTToken(t string) (*jwt.Token, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unauth")
		}
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
