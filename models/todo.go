package models

import "bubble/dao"

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}
func GetAllTodo() (todolist []*Todo, err error) {
	err = dao.DB.Find(&todolist).Error
	return
}

func UpdateATodo
