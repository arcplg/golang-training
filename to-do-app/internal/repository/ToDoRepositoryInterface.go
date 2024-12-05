package repository

import (
	model "to-do-app/internal/models"
)

type ToDoRepository interface {
	GetAll() ([]model.ToDo, error)
	Create(todo model.ToDo) (int64, error)
	GetDetailById(id int) (*model.ToDo, error)
	Update(toDo *model.ToDo) error
	Delete(id int) error
}
