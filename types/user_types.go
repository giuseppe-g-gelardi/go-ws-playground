package types

type User struct {
	Id        string `bson:"_id,omitempty"` // primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string `bson:"first_name,omitempty"`
	LastName  string `bson:"last_name,omitempty"`
	Email     string `bson:"email,omitempty"`
}

