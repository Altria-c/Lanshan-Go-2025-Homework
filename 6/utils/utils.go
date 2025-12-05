package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var customsecret = []byte("南风知我意吹梦到西洲")

type Myclaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 生成token
func MakeToken(username string) (string, error) {
	claim := Myclaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
			Issuer:    "yang",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(customsecret)
}

// 生成refreshtoken
func MakeRefreshToken(username string) (string, error) {
	claim := Myclaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			Issuer:    "yang",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(customsecret)
}

// 解析token
func ParseToken(tokenString string) (*Myclaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Myclaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return customsecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 类型断言
	if clames, ok := token.Claims.(*Myclaims); ok && token.Valid {
		return clames, nil
	}
	return nil, errors.New("invalid token")
}
