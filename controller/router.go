package controller

// import (
// 	"github.com/gin-gonic/gin"
// )

// // 外部パッケージに公開するインタフェース
// type Router interface {
// 	HandleUsersRequest(ctx *gin.Context)
// }

// type router struct {
// 	uc UserController
// }

// func NewRouter(uc UserController) Router {
// 	return &router{uc: uc}
// }

// func (ro *router) HandleUsersRequest(ctx *gin.Context) {
// 	switch ctx.Request.Method {
// 	case "POST":
// 		ro.uc.SignUp(ctx)
// 	case "GET":
// 		ro.uc.SelectUser(ctx)
// 	case "PUT":
// 		ro.uc.UpdateUser(ctx)
// 	case "DELETE":
// 		ro.uc.DeleteUser(ctx)
// 	}
// }
