package controllers_helpers

import (
	"album-ranking-user-albums/spotify_models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUserProfile(token string) (*spotify_models.Profile, error) {

	url := "https://api.spotify.com/v1/me"

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

	var result spotify_models.Profile

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
