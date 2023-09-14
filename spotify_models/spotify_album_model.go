package spotify_models

type ID string

type URI string

type Image struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	URL    string `json:"url"`
}

type SimpleArtist struct {
	Name         string            `json:"name"`
	ID           ID                `json:"id"`
	URI          URI               `json:"uri"`
	Endpoint     string            `json:"href"`
	ExternalURLs map[string]string `json:"external_urls"`
}

type SimpleAlbum struct {
	Name                 string            `json:"name"`
	Artists              []SimpleArtist    `json:"artists"`
	AlbumGroup           string            `json:"album_group"`
	AlbumType            string            `json:"album_type"`
	ID                   ID                `json:"id"`
	URI                  URI               `json:"uri"`
	AvailableMarkets     []string          `json:"available_markets"`
	Endpoint             string            `json:"href"`
	Images               []Image           `json:"images"`
	ExternalURLs         map[string]string `json:"external_urls"`
	ReleaseDate          string            `json:"release_date"`
	ReleaseDatePrecision string            `json:"release_date_precision"`
}

type Copyright struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type FullAlbum struct {
	SimpleAlbum
	Copyrights  []Copyright       `json:"copyrights"`
	Genres      []string          `json:"genres"`
	Popularity  int               `json:"popularity"`
	Tracks      SimpleTrackPage   `json:"tracks"`
	ExternalIDs map[string]string `json:"external_ids"`
}

type SavedAlbum struct {
	AddedAt   string `json:"added_at"`
	FullAlbum `json:"album"`
}

type SimpleTrackPage struct {
	basePage
	Tracks []SimpleTrack `json:"items"`
}

type basePage struct {
	Endpoint string `json:"href"`
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
	Total    int    `json:"total"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type SimpleTrack struct {
	Album            SimpleAlbum       `json:"album"`
	Artists          []SimpleArtist    `json:"artists"`
	AvailableMarkets []string          `json:"available_markets"`
	DiscNumber       int               `json:"disc_number"`
	Duration         int               `json:"duration_ms"`
	Explicit         bool              `json:"explicit"`
	ExternalURLs     map[string]string `json:"external_urls"`
	ExternalIDs      TrackExternalIDs  `json:"external_ids"`
	Endpoint         string            `json:"href"`
	ID               ID                `json:"id"`
	Name             string            `json:"name"`
	PreviewURL       string            `json:"preview_url"`
	TrackNumber      int               `json:"track_number"`
	URI              URI               `json:"uri"`
	Type             string            `json:"type"`
}

type TrackExternalIDs struct {
	ISRC string `json:"isrc"`
	EAN  string `json:"ean"`
	UPC  string `json:"upc"`
}

type SavedAlbumPage struct {
	basePage
	Albums []SavedAlbum `json:"items"`
}
