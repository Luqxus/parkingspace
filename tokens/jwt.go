package tokens

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)


var key string = "ioeur09bc80934r8c0b80e8r09bfc89043runmewifd-0c[wen8r9po"

type Claims struct {
	 UID string
	 Email string
	 jwt.RegisteredClaims
}

func GenerateJWT(uid, email string) (string, error) {
	claims := &Claims{
		UID: uid,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Hour).Local()),
		},
	}
	
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(key))
}

func VerifyToken(signedToken string) (string, error) {
	claims := new(Claims)
	
	token, err := jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return "", err
	}
	
	if !token.Valid {
		return "", errors.New("invalid authorization header")
	}
	
	return claims.UID, nil
}