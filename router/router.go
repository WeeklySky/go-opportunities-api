package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// initialize router using gin default engine configs
	r := gin.Default()

	// initialize routes
	initializeRoutes(r)
	//run the server
	r.Run() // listen and serve on 0.0.0.0:8080
}
