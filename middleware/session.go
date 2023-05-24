package middleware

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func SessionCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// クッキーからトークンを取得
		token, err := ctx.Cookie("SessionID")

		if err != nil {
			ctx.JSON(401, gin.H{"message": "Unauthorized"})
			ctx.Abort()
			return
		}

		// トークンのクレームを取得
		claims := jwt.MapClaims{}
		_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil {
			fmt.Print(err)
			ctx.JSON(401, gin.H{"message": "Unauthorized"})
			ctx.Abort()
			return
		}
		if claims["user_id"] == nil {
			ctx.JSON(401, gin.H{"message": "Unauthorized"})
			ctx.Abort()
			return
		}

		// contextにユーザーIDをセット
		ctx.Keys["user_id"] = claims["user_id"]
		ctx.Next()
	}
}
