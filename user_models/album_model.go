package user_models

type Album struct {
	Id          string   `bson:"id"`
	Images      []string `bson:"images"`
	Name        string   `bson:"name"`
	ReleaseDate string   `bson:"release_date"`
	ArtistName  string   `bson:"artist_name"`
	Popularity  int      `bson:"popularity"`
	Score       int      `bson:"score"`
	AddedAt     string   `bson:"added_at"`
}

type NewScore struct {
	Score int `json:"score"`
}

type NestedAlbum struct {
	Albums Album `bson:"albums"`
}
