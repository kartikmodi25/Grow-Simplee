package database

import (
	"backend-assignment/responses"
	"context"
	"database/sql"
)

type DB interface {
	BeginTxn(opts ...*sql.TxOptions) DB
	Commit() error
	Rollback() error

	AutoMigrate() error
	CheckExistingUser(ctx context.Context, email string) (bool, error)
	CheckUserCredentials(ctx context.Context, email string, password string) (bool, error)
	CreateUser(ctx context.Context, name string, email string, password string) error
	UpdateMovieRating(ctx context.Context, movieName string, rating int8) (float32, error)
	GetMoviesData(ctx context.Context) ([]string, error)
	GetMovieRatings(ctx context.Context) ([]responses.MovieRating, error)
}
