package routes

import (
	"album-ranking-user-albums/controllers"
	"github.com/gin-gonic/gin"
)

func SessionRoute(router *gin.Engine) {
	router.GET("/update/:token", controllers.UpdateAlbums())
	router.GET("/albums/year/:year/:user_id", controllers.AlbumsByYear())
	router.GET("/albums/decade/:decade/:user_id", controllers.AlbumsByDecade())
}
