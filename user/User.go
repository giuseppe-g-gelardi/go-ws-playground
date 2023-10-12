package user

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

func (u *User) NewUser(userName string, email string) *User {

	fmt.Println("Creating a new user")

	user := &User{
		Id:       uuid.New().String(),
		UserName: userName,
		Email:    email,
	}

	fmt.Printf("New user created: \n")
	fmt.Printf("User id: %s\n", user.Id)
	fmt.Printf("User name: %s\n", user.UserName)
	fmt.Printf("Email: %s\n", user.Email)

	return user
}
