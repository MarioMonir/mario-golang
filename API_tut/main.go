package main

import (
	"mario/src/Configs/Database"
	"mario/src/Entities/User"

	"github.com/gin-gonic/gin"
)

// ==========================================

func main() {

	Database.DbConnectAndMigrate()

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~

	app := gin.Default()
	v1 := app.Group("/api")

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~

	User.UserRouter(v1.Group("/user"))

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~

	app.Run(":5000")
}
