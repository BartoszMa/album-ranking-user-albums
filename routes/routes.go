package routes

import (
	"album-ranking-user-albums/controllers"
	"github.com/gin-gonic/gin"
)

func SessionRoute(router *gin.Engine) {
	router.GET("/update/:token", controllers.UpdateAlbums())
}
