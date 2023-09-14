package user_models

type User struct {
	Id     string  `bson:"id"`
	Albums []Album `bson:"albums"`
}
