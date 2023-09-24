package controllers

import (
	"album-ranking-user-albums/responses"
	"album-ranking-user-albums/user_models"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AlbumSetScore() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var newScore user_models.NewScore

		if err := c.BindJSON(&newScore); err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error date parse", Data: map[string]interface{}{"error": err.Error()}})
			return
		}

	}
}
