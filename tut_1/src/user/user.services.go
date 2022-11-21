package user

// - - - - - - - - - - - - - - - - - -

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// - - - - - - - - - - - - - - - - - -

func GetUserService(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"user": "mario"})
}
