package User

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ========================================================

func userListController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "USER GET LIST",
		"data":    userGetListService(),
	})
}

// ========================================================

func userGetOneController(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "USER GET ONE",
		"data":    userGetOneService(id),
	})
}

// ========================================================

func userCreateController(c *gin.Context) {
	// user := c.Request.Body
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "USER CREATED",
		// "data":    UserCreateService(user),
	})
}

// ========================================================

func userUpdateController(c *gin.Context) {
	// id, _ := strconv.Atoi(c.Param("id"))
	// user := c.Request.Body
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "USER UPDATED",
		// "data":    UserUpdateService(id, user),
	})
}

// ========================================================

func userDeleteController(c *gin.Context) {
	// id, _ := strconv.Atoi(c.Param("id"))
	// user := c.Request.Body
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "USER DELETED",
		// "data":    UserUpdateService(id, user),
	})
}

// ========================================================
