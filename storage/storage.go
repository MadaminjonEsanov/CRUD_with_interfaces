package storage

import (
	"app/storage/postgres"
	repoi "app/storage/repoI"

	"github.com/jackc/pgx/v5"
)

type StorageI interface {
	GetUserRepo() repoi.UserRepoI
	GetTodoRepo() repoi.TodoRepoI
}

type Storage struct {
	UserRepo repoi.UserRepoI
	TodoRepo repoi.TodoRepoI
}

func NewStorage(conn *pgx.Conn) StorageI {
	storage := &Storage{
		UserRepo: postgres.NewUserRepo(conn),
		TodoRepo: postgres.NewTodoRepo(conn),
	}

	return storage
}

func (s *Storage) GetUserRepo() repoi.UserRepoI {
	return s.UserRepo
}

func (s *Storage) GetTodoRepo() repoi.TodoRepoI {
	return s.TodoRepo
}