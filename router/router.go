package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// inicializa o Router utilizando as configs default do gin
	r := gin.Default()
	// define uma rota
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
