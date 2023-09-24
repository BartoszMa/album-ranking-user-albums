package controllers

import (
	"album-ranking-user-albums/responses"
	"album-ranking-user-albums/user_models"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strconv"
	"time"
)

func AlbumsByDecade() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

		decade := c.Param("decade")
		userId := c.Param("user_id")
		defer cancel()

		yr, err := strconv.Atoi(decade)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error date parse", Data: map[string]interface{}{"error": err.Error()}})
			return
		}

		pipeline := []bson.M{
			{
				"$match": bson.M{
					"id": userId,
				},
			},
			{
				"$unwind": "$albums",
			},
			{
				"$match": bson.M{
					"albums.release_date": bson.M{
						"$gte": strconv.Itoa(yr),
						"$lte": strconv.Itoa(yr + 10),
					},
				},
			},
			{
				"$project": bson.M{
					"_id": 0, "albums": 1,
				},
			},
			{
				"$replaceRoot": bson.M{"newRoot": "$albums"},
			},
		}

		var results []user_models.Album

		cursor, err := userCollection.Aggregate(ctx, pipeline)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error date parse", Data: map[string]interface{}{"error": err.Error()}})
			return
		}

		if err := cursor.All(ctx, &results); err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error date parse", Data: map[string]interface{}{"error": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "ok", Data: map[string]interface{}{"data": results}})

	}
}
