package User

import "gorm.io/gorm"

// ---------------------------------------

type User struct {
	gorm.Model
	Id    int    `json:"id"    gorm:"primary_key"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}

// ---------------------------------------

// var Users = []User{
// 	{
// 		Id:    1,
// 		Name:  "Mario",
// 		Age:   24,
// 		Phone: "01201200553",
// 	},
// 	{
// 		Id:    2,
// 		Name:  "Maivel",
// 		Age:   16,
// 		Phone: "01228454015",
// 	},
// }
