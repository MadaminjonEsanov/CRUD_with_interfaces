package postgres

import (
	"app/models"
	repoi "app/storage/repoI"
	"context"

	"github.com/jackc/pgx/v5"
)

type UserRepo struct {
	conn *pgx.Conn
}

func NewUserRepo(conn *pgx.Conn) repoi.UserRepoI {

	userRepo := &UserRepo{
		conn: conn,
	}

	return userRepo
}

func (u *UserRepo) CreateUser(ctx context.Context, user models.User) error {
	return nil
}

func (u *UserRepo) GetUsers(ctx context.Context, limit, page string) ([]models.User, error) {
	return nil, nil
}

func (u *UserRepo) GetUserById(ctx context.Context, userId string) (*models.User, error) {
	return nil, nil
}

func (u *UserRepo) UpdateUser(ctx context.Context, user models.User) error {
	return nil
}

func (u *UserRepo) DeleteUserById(ctx context.Context, userId string) error {
	return nil
}