package main

// - - - - - - - - - - - - - - - - - - - -

import (
	"mario/src/entities/user"

	"github.com/gin-gonic/gin"
)

// - - - - - - - - - - - - - - - - - - - -

func main() {
	app := gin.Default()
	v1 := app.Group("/api")

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~

	user.UserRouter(v1.Group("/user"))

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~

	app.Run(":5000")
}
