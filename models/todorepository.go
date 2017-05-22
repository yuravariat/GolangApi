package models

import (
	"GoApi/entities"
	"fmt"
)

var currentId int

var todos entities.Todos

// Give us some seed data
func init() {
	TodoRepoCreateTodo(entities.Todo{Name: "Write presentation"})
	TodoRepoCreateTodo(entities.Todo{Name: "Host meetup"})
}

func TodoRepoGetAll() []entities.Todo {
	if todos != nil {
		return todos
	}
	// return empty Todos if not found
	return entities.Todos{}
}

func TodoRepoFindTodo(id int) entities.Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return entities.Todo{}
}

//this is bad, I don't think it passes race condtions
func TodoRepoCreateTodo(t entities.Todo) entities.Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func TodoRepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
