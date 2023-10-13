package mocks

import (
	"playground.com/m/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User types.User

var Users = []User{
	{Id: primitive.NewObjectID(), FirstName: "John", LastName: "Doe", Email: "johndoe@email.com"},
	{Id: primitive.NewObjectID(), FirstName: "Jane", LastName: "Doe", Email: "janedoe@email.com"},
	{Id: primitive.NewObjectID(), FirstName: "Jack", LastName: "Doe", Email: "jackdoe@email.com"},
	{Id: primitive.NewObjectID(), FirstName: "Jill", LastName: "Doe", Email: "jilldoe@email.com"},
	{Id: primitive.NewObjectID(), FirstName: "James", LastName: "Doe", Email: "jamesdoe@emai.com"},
	{Id: primitive.NewObjectID(), FirstName: "Jenny", LastName: "Doe", Email: "jennydoe@email.com"},
}

func GetMockUsers() []User {
	return Users
}
