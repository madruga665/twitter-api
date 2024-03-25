package main

import (
	"github.com/gin-gonic/gin"
	"github.com/madruga665/twitter-api/api/routes"
)

func main() {
	app := gin.Default()

	routes.Routes(app)

	app.Run("localhost:3001")
}
