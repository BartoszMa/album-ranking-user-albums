package controllers

import (
	"album-ranking-user-albums/controllers_helpers"
	"album-ranking-user-albums/dbDriver"
	"album-ranking-user-albums/spotify_models"
	"album-ranking-user-albums/user_models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var userCollection = dbDriver.GetCollection(dbDriver.DB, "users")

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
				Images:      []string{album.Images[0].URL, album.Images[1].URL, album.Images[2].URL},
				Name:        album.Name,
				ReleaseDate: album.ReleaseDate,
				ArtistName:  album.Artists[0].Name,
				Popularity:  album.Popularity,
				Score:       0,
			}

			user.Albums = append(user.Albums, newAlbum)
		}
	}

	if result.Next == "" {
		return user, nil
	}

	return fetchAlbums(token, result.Next, user)
}

//func CheckIfUserExist(id string) {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	var user user_models.User
//	defer cancel()
//	err := userCollection.FindOne(ctx, bson.M{"id": id}).Decode(&user)
//}

func UpdateAlbums() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Param("token")

		//url := "https://api.spotify.com/v1/me/albums?limit=50"
		//
		//user := user_models.User{
		//	Id: "user_id",
		//}

		//response, err := fetchAlbums(token, url, &user)
		//if err != nil {
		//	c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error fetching albums", Data: map[string]interface{}{"error": err.Error()}})
		//	return
		//}
		//
		//fmt.Println(response.Albums[0])

		fmt.Println("")

		profile, err := controllers_helpers.GetUserProfile(token)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(profile.ID)

		//c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "ok", Data: map[string]interface{}{"data": "ok"}})

	}
}
