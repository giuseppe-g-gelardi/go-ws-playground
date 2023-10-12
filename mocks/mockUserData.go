package mocks

import (
	"playground.com/m/types"

	"github.com/google/uuid"
)

type User types.User

var Users = []User{
	{Id: uuid.New(), FirstName: "John", LastName: "Doe", Email: "johndoe@email.com"},
	{Id: uuid.New(), FirstName: "Jane", LastName: "Doe", Email: "janedoe@email.com"},
	{Id: uuid.New(), FirstName: "Jack", LastName: "Doe", Email: "jackdoe@email.com"},
	{Id: uuid.New(), FirstName: "Jill", LastName: "Doe", Email: "jilldoe@email.com"},
	{Id: uuid.New(), FirstName: "James", LastName: "Doe", Email: "jamesdoe@emai.com"},
	{Id: uuid.New(), FirstName: "Jenny", LastName: "Doe", Email: "jennydoe@email.com"},
}

func GetMockUsers() []User {
	return Users
}
