package repoi

import (
    "context"
    "app/models"
)

type UserRepoI interface {

	CreateUser(ctx context.Context, user models.User)  error
	GetUserByUsername(ctx context.Context, limit,page string) ( []models.User, error)
	GetUserById(ctx context.Context, userId string) ( models.User, error)
	UpdateUser(ctx context.Context, user models.User)  error
	DeleteUserById(ctx context.Context, userId string) error
}


type TodoRepoI interface {

	CreateTodo(ctx context.Context, Todo models.Todo)  error
	GetTodos(ctx context.Context, limit,page string) ( []models.GetTodosResponse, error)
	GetTodoById(ctx context.Context, TodoId string) ( models.Todo, error)
	UpdateTodo(ctx context.Context, Todo models.Todo)  error
	DeletetodoById(ctx context.Context, TodoId string) error
}