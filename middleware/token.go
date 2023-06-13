package middleware

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// クッキーからトークンを取得
		cookie, err := ctx.Cookie("AccessToken")
		if err != nil {
			http.Error(ctx.Writer, "Unauthorized", http.StatusUnauthorized)
			ctx.Abort()
			return
		}

		// トークンの検証
		tokenString := cookie
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// 秘密鍵のバイト配列を返す
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil || !token.Valid {
			http.Error(ctx.Writer, "Unauthorized", http.StatusUnauthorized)
			ctx.Abort()
			return
		}

		// トークンの検証が成功した場合、次のハンドラを呼び出す
		ctx.Next()
	}
}
