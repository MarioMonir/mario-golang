package user

// - - - - - - - - - - - - - - - - - - - - - - - - - -

import (
	"github.com/gin-gonic/gin"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - -

func UserRouter(userRouter *gin.RouterGroup) {
	userRouter.GET("/", UserListController)
	userRouter.GET("/:id", UserGetOneController)
}
