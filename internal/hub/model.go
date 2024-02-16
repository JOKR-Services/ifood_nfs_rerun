package hub

type Hub struct {
	From string `bson:"from"`
	To   string `bson:"to"`
}
