package controllers

import (
	"album-ranking-user-albums/controllers_helpers"
	"album-ranking-user-albums/dbDriver"
	"album-ranking-user-albums/responses"
	"album-ranking-user-albums/spotify_models"
	"album-ranking-user-albums/user_models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

var userCollection = dbDriver.GetCollection(dbDriver.DB, "albums")

func fetchAlbums(token string, url string, user *user_models.User) (*user_models.User, error) {

	bearer := "Bearer " + token

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", bearer)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code %d", res.StatusCode)
	}

	var result spotify_models.SavedAlbumPage

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	if len(result.Albums) > 0 {
		for _, album := range result.Albums {
			newAlbum := user_models.Album{
				Id:          album.ID,
				Images:      []string{album.Images[0].URL, album.Images[1].URL, album.Images[2].URL},
				Name:        album.Name,
				ReleaseDate: album.ReleaseDate,
				ArtistName:  album.Artists[0].Name,
				Popularity:  album.Popularity,
				Score:       0,
				AddedAt:     album.AddedAt,
			}

			user.Albums = append(user.Albums, newAlbum)
		}
	}

	if result.Next == "" {
		return user, nil
	}

	return fetchAlbums(token, result.Next, user)
}

func getAlbumAndCompareWithDate(token, url string, date time.Time, array []user_models.Album) ([]user_models.Album, error) {
	bearer := "Bearer " + token

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", bearer)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code %d", res.StatusCode)
	}

	var result spotify_models.SavedAlbumPage

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	addedAtDate, err := time.Parse("2006-01-02T15:04:05Z", result.Albums[0].AddedAt)
	if err != nil {
		return nil, err
	}

	if addedAtDate.After(date) {

		if len(result.Albums) > 0 {
			for _, album := range result.Albums {
				newAlbum := user_models.Album{
					Id:          album.ID,
					Images:      []string{album.Images[0].URL, album.Images[1].URL, album.Images[2].URL},
					Name:        album.Name,
					ReleaseDate: album.ReleaseDate,
					ArtistName:  album.Artists[0].Name,
					Popularity:  album.Popularity,
					Score:       0,
					AddedAt:     album.AddedAt,
				}

				array = append(array, newAlbum)
			}
		}

		return getAlbumAndCompareWithDate(token, result.Next, date, array)
	}

	return array, nil
}

func UpdateAlbums() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

		token := c.Param("token")
		var user user_models.User

		defer cancel()

		url := "https://api.spotify.com/v1/me/albums?limit=50"

		profile, err := controllers_helpers.GetUserProfile(token)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error fetching albums",
				Data:    map[string]interface{}{"error": err.Error()}})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"id": profile.ID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error getting user",
				Data:    map[string]interface{}{"error": err.Error()}})
			return
		}
		if count == 0 {
			response, err := fetchAlbums(token, url, &user)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{
					Status:  http.StatusInternalServerError,
					Message: "error fetching albums",
					Data:    map[string]interface{}{"error": err.Error()}})
				return
			}
			newUser := user_models.User{Id: profile.ID, Albums: response.Albums}
			_, err = userCollection.InsertOne(ctx, newUser)
		} else {
			showLoadedCursor, err := userCollection.Aggregate(ctx,
				[]bson.M{{"$match": bson.M{"id": profile.ID}}, {"$unwind": "$albums"}, {"$sort": bson.M{"albums.added_at": -1}}, {"$limit": 1}, {"$project": bson.M{"_id": 0, "albums": 1}}},
			)

			var showsLoaded []user_models.NestedAlbum
			if err = showLoadedCursor.All(ctx, &showsLoaded); err != nil {
				panic(err)
			}

			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error fetching user data", Data: map[string]interface{}{"error": err.Error()}})
				return
			}

			url = "https://api.spotify.com/v1/me/albums?limit=1"

			addedAtDate, err := time.Parse("2006-01-02T15:04:05Z", showsLoaded[0].Albums.AddedAt)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error fetching user data", Data: map[string]interface{}{"error": err.Error()}})
				return
			}

			var albums []user_models.Album

			newAlbums, err := getAlbumAndCompareWithDate(token, url, addedAtDate, albums)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error fetching user data", Data: map[string]interface{}{"error": err.Error()}})
				return
			}

			for _, album := range newAlbums {
				_, err = userCollection.UpdateOne(ctx, bson.M{"id": profile.ID}, bson.M{"$push": bson.M{"albums": album}})
				fmt.Println(album)
				if err != nil {
					c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error fetching user data", Data: map[string]interface{}{"error": err.Error()}})
					return
				}
			}

		}

		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "ok", Data: map[string]interface{}{"data": profile.ID}})
		return

	}
}
