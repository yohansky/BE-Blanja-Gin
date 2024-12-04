package main

import (
	"be-blanja/src/config"
	"be-blanja/src/helper"
	"be-blanja/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// gotenv.Load()
	config.InitDB()
	helper.Migrate()
	app := gin.New()

	routes.Router(app)

	app.Run(":8080")
}
