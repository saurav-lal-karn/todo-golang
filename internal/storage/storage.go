package storage

import "time"

type Todo struct {
	Id        string
	Item      string
	CreatedAt time.Time
	DueDate   time.Time
	Completed string
}

type Storage interface {
	init() error
	AddTodo(todo *Todo) error
	List() ([]*Todo, error)
	Update(id string, todo *Todo) error
	Delete(id string) error
	Complete(id string) error
}
