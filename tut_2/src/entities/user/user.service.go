package user

import (
	"fmt"

	"golang.org/x/exp/slices"
)

// - - - - - - - - - - - - - - - - - -

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}

// - - - - - - - - - - - - - - - - - -

// Dummy Data
var Users = []User{
	{
		Id:    1,
		Name:  "Mario",
		Age:   24,
		Phone: "01201200553",
	},
	{
		Id:    2,
		Name:  "Maivel",
		Age:   16,
		Phone: "01228454015",
	},
}

// - - - - - - - - - - - - - - - - - -

func UserGetListService() []User {
	return Users
}

// - - - - - - - - - - - - - - - - - -

func UserGetOneService(id int) User {

	fmt.Println("ID : ", id)
	idx := slices.IndexFunc(Users, func(user User) bool { return user.Id == id })

	return Users[idx]
}
