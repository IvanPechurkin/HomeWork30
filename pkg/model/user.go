package model

import (
	"errors"
)

type User struct {
	Id      int    `json:"id" bson:"_id"`
	Name    string `json:"name" bson:"name"`
	Age     int    `json:"age" bson:"age"`
	Friends []int  `json:"friends" bson:"friends"`
}

// Добавление друга
func (u *User) AddFriend(userId int) error {
	index := findIndex(u.Friends, userId)
	if index == -1 {
		u.Friends = append(u.Friends, userId)
	} else {
		return errors.New("У пользователя " + u.Name + " уже есть такой друг")
	}
	return nil
}

// Удаление друга
func (u *User) DeleteFriend(userId int) {
	if len(u.Friends) < 1 {
		return
	}
	index := findIndex(u.Friends, userId)
	if index > -1 {
		u.Friends = append(u.Friends[:index], u.Friends[index+1:]...)
	}
}

// Поиск индекса в слайсе
func findIndex(list []int, userId int) int {
	for i := 0; i < len(list); i++ {
		if list[i] == userId {
			return i
		}
	}
	return -1
}
