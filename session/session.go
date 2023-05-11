package session

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionInfo struct {
	UserId interface{}
}

func SessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		var sessionInfo SessionInfo
		session := sessions.Default(c)
		sessionInfo.UserId = session.Get("UserId")

		// セッションがない場合、ログインフォームをだす
		if sessionInfo.UserId == nil {
			log.Println("ログインしていません")
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
		} else {
			c.Set("UserId", sessionInfo.UserId)
			c.Next()
		}
		log.Println("ログインチェック終わり")
	}
}
