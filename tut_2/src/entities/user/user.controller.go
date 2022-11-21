package user

// - - - - - - - - - - - - - - - - - - - - - - - - - -

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - -

func UserListController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "USER GET LIST",
		"data":    UserGetListService(),
	})
}

// - - - - - - - - - - - - - - - - - - - - - - - - - -

func UserGetOneController(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "USER GET ONE",
		"data":    UserGetOneService(id),
	})
}

// - - - - - - - - - - - - - - - - - - - - - - - - - -
