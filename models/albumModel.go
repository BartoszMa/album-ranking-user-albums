package models

type Album struct {
	id          string    `bson:"id"`
	images      [3]string `bson:"images"`
	name        string    `bson:"name"`
	releaseDate string    `bson:"release_date"`
	artistName  string    `bson:"artistName"`
	popularity  int       `bson:"popularity"`
	score       int       `bson:"score"`
}
