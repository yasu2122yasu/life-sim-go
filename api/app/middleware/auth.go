package middleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secret") // 本番環境ではもっと安全な方法でキーを管理する

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// tokenの発行はフロントエンドで行い、クレームの設定と署名をバックエンドで行う
func GenerateToken() (string, error) {
	// JWTトークンの生成
	expirationTime := time.Now().Add(24 * time.Hour)
	// claimsの設定
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}
	// 署名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil // CreateTokenにて指定した文字列を使います
	})
	fmt.Println(token)
	if err != nil {
		return nil, err
	}
	return token, nil
}
