package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	// Iss string `json:"iss"`
	Sub string `json:"sub"`
	Exp int64  `json:"exp"`
}

func GenerateToken(userID string) (string, error) {
	key := os.Getenv("SECRET_KEY")
	// JWTの署名に使用するシークレットキー
	secretKey := []byte(key)

	// トークンの有効期限を設定
	expirationTime := time.Now().Add(1 * time.Hour * 24 * 30)
	// expirationTime := time.Now().Add(time.Second * 10)

	// トークンに埋め込むクレーム（情報）
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": expirationTime.Unix(),
	}

	// トークンを作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークンに署名を付けて文字列に変換
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
