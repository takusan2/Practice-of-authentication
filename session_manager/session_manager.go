package session_manager

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/takuya-okada-01/heart-note/infrastructure/redis_client"
)

type SessionManager interface {
	SessionCheck() gin.HandlerFunc
	SetSession(c *gin.Context) error
}

type sessionManager struct {
	redisClient redis_client.RedisClient
}

func NewSessionManager(redisClient redis_client.RedisClient) SessionManager {
	return &sessionManager{redisClient: redisClient}
}

func (s *sessionManager) SessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		var sessionInfo SessionInfo
		session := sessions.Default(c)
		sessionInfo.SessionID = session.Get("SessionID")
		sessionInfo.UserID = session.Get("UserID")

		// セッションがない場合、ログインフォームをだす
		if sessionInfo.SessionID == nil {
			log.Println("ログインしていません")
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
		} else {
			log.Println("ログインチェック中")
			sessionInfo.SessionID = uuid.New().String()
			// セッションIDをcookieにセット
			c.Set("SessionID", sessionInfo.SessionID)
			// セッションIDをredisにセット
			s.redisClient.SetSession(c, sessionInfo.SessionID.(string), sessionInfo.UserID, 60*60*24*30)
			c.Next()
		}
		log.Println("ログインチェック終わり")
	}
}

func (s *sessionManager) SetSession(c *gin.Context) error {
	session := sessions.Default(c)
	var sessionInfo SessionInfo
	sessionInfo.SessionID = uuid.New().String()
	// セッションIDをcookieにセット
	session.Set("SessionID", sessionInfo.SessionID)
	// セッションIDをredisにセット
	s.redisClient.SetSession(c, sessionInfo.SessionID.(string), sessionInfo.UserID, 60*60*24*30)
	return session.Save()
}
