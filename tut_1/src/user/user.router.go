package user

// - - - - - - - - - - - - - - - - - -

import "github.com/gin-gonic/gin"

// - - - - - - - - - - - - - - - - - -

func UsersRouter(router *gin.RouterGroup) {
	router.GET("/", GetUserService)
}
