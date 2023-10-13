package types

type User struct {
	// Id        primitive.ObjectID `bson:"_id,omitempty"`
	Id        string `bson:"_id,omitempty"`
	FirstName string `bson:"first_name,omitempty"`
	LastName  string `bson:"last_name,omitempty"`
	Email     string `bson:"email,omitempty"`
}

