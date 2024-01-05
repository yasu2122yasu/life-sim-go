package middleware

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// TODO: 本番環境では安全な方法でキーを生成し、管理する
var jwtKey = []byte("your-very-strong-secret-key") // 安全なキーをここで設定

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// トークンの発行。フロントエンドからクレーム（例：ユーザーのメール）を受け取り、トークンに含める
func GenerateToken(email string) (string, error) {
	// トークンの有効期限を設定
	expirationTime := time.Now().Add(24 * time.Hour)
	// クレームの設定
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
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
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 署名方法の検証
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		// jwtKeyを使用してトークンを検証
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	// キャストしてクレームを検証
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		fmt.Printf("User Email: %v\n", claims.Email)
		return token, nil
	} else {
		return nil, errors.New("Invalid token")
	}
}
