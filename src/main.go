package main

import (
	"os"

	"github.com/WinIT23/web-service-gin/configs"
	"github.com/WinIT23/web-service-gin/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	// Setup MongoDB.
	configs.GetMongoClient()

	// Setup routers.
	routers.AlbumRoute(app)

	// Start server.
	app.Run(getHost())
}

func getHost() string {
	host := os.Getenv("GO_RUN_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("GO_RUN_PORT")
	if port == "" {
		port = "8080"
	}
	return host + ":" + port
}
