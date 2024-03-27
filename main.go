package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/madruga665/twitter-api/api/routes"
	connection "github.com/madruga665/twitter-api/db"
)

func main() {
	error := connection.OpenDatabase()

	if error != nil {
		log.Printf("eu ruim na hora de conectar no db erro: %v", error)
	}

	defer connection.CloseDatabase()

	app := gin.Default()

	routes.Routes(app)

	app.Run("localhost:3001")
}
