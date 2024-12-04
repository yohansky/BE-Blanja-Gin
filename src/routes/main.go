package routes

import "github.com/gin-gonic/gin"

func Router(app *gin.Engine) {
	app.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"Message": "Hello world"})
	})
}
