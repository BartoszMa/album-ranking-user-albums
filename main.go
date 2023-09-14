package main

import (
	"album-ranking-user-albums/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {

	router := gin.Default()

	routes.SessionRoute(router)

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("ALBUM_RANKING_API_GATEWAY")},
		AllowMethods: []string{"POST", "GET", "DELETE"},
	}))

	errRouter := router.Run(os.Getenv("ALBUM_RANKING_USER_ALBUMS_URI"))
	if errRouter != nil {
		panic(errRouter)
	}
}
