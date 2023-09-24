package spotify_models

type Profile struct {
	DisplayName  string            `json:"display_name"`
	ExternalURLs map[string]string `json:"external_urls"`
	Followers    Followers         `json:"followers"`
	Endpoint     string            `json:"href"`
	ID           string            `json:"id"`
	Images       []Image           `json:"images"`
	URI          string            `json:"uri"`
	Country      string            `json:"country"`
	Email        string            `json:"email"`
	Product      string            `json:"product"`
}

type Followers struct {
	Count    uint   `json:"total"`
	Endpoint string `json:"href"`
}
