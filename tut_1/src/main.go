package main

// ==================================================

import (
	"github.com/gin-gonic/gin"
	// "Entity/User/User.router.go"
)

// ==================================================

func main() {
	app := gin.Default()

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~

	// v1 := app.Group("")
	// {
	// 	v1.GET("/", IndexRouter)
	// }

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~

	app.Run(":5000")

}
