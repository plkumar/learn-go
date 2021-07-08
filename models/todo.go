package models

import (
	"errors"
	"fmt"
)

type Todo struct {
	ID        int
	Text      string
	Completed bool
}

var (
	todos   []*Todo
	ID_SEED int = 1
)

func All() []*Todo {
	return todos
}

func Add(todo Todo) (Todo, error) {
	if todo.ID != 0 {
		return Todo{}, errors.New("invalid user passed, id should not be specified")
	}
	todo.ID = ID_SEED
	ID_SEED++
	todos = append(todos, &todo)
	return todo, nil
}

func Update(todo Todo) (Todo, error) {
	for i, v := range todos {
		if v.ID == todo.ID {
			todos[i] = &todo
			return todo, nil
		}
	}
	return Todo{}, fmt.Errorf("Todo with id : %d not found", todo.ID)
}

func RemoveByID(id int) error {
	if len(todos) > 0 {
		for i, v := range todos {
			if v.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				return nil
			}
		}
		return errors.New("todo not found with id specified")
	}
	return nil
}
