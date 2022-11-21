package User

import "github.com/gin-gonic/gin"

func UserRouter(router *gin.RouterGroup) {
	router.GET("/", userListController)
	router.POST("/", userCreateController)
	router.GET("/:id", userGetOneController)
	router.PUT("/:id", userUpdateController)
	router.DELETE("/:id", userDeleteController)
}
