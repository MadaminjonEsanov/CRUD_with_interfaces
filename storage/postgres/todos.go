package postgres

import (
	"app/models"
	repoi "app/storage/repoI"
	"context"

	"github.com/jackc/pgx/v5"
)

type TodoRepo struct {
	conn *pgx.Conn
}

func NewTodoRepo(conn *pgx.Conn) repoi.TodoRepoI {
	todoRepo := &TodoRepo{conn: conn}

	return todoRepo
}

func (t *TodoRepo) CreateTodo(ctx context.Context, todo models.Todo) error {
	return nil
}

func (t *TodoRepo) GetTodos(ctx context.Context, limit, page string) (*models.GetTodosResponse, error) {
	return nil, nil
}

func (t *TodoRepo) GetTodoById(ctx context.Context, todoId string) (*models.Todo, error) {
	return nil, nil
}

func (t *TodoRepo) UpdateTodo(ctx context.Context, todo models.Todo) error {
	return nil
}

func (t *TodoRepo) DeleteTodoById(ctx context.Context, todoId string) error {
	return nil
}