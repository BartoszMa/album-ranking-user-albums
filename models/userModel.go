package models

type User struct {
	id     string  `bson:"id"`
	albums []Album `bson:"albums"`
}
