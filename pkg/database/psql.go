package database

import (
	"app/config"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)



var ctx =context.Background()

func Conn(cfg config.Config) (*pgx.Conn, error){

	dbURL:= fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",

		cfg.DbUser,
		cfg.DbPassword,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName,
		
	) 

	return pgx.Connect(ctx, dbURL)
}