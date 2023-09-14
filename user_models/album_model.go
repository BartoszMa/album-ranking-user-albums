package user_models

type Album struct {
	Id          string   `bson:"id"`
	Images      []string `bson:"images"`
	Name        string   `bson:"name"`
	ReleaseDate string   `bson:"release_date"`
	ArtistName  string   `bson:"artistName"`
	Popularity  int      `bson:"popularity"`
	Score       int      `bson:"score"`
}
