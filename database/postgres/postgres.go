package postgres

import (
	"backend-assignment/config"
	"backend-assignment/database"
	"backend-assignment/database/models"
	"backend-assignment/responses"
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	pq "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgres struct {
	conf config.Database
	_db  *gorm.DB
}

const DSN = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Kolkata"

func (pg *postgres) BeginTxn(opts ...*sql.TxOptions) database.DB {
	return &postgres{
		conf: pg.conf,
		_db:  pg._db.Begin(opts...),
	}
}

func (pg *postgres) Commit() error {
	return pg._db.Commit().Error
}

func (pg *postgres) Rollback() error {
	return pg._db.Rollback().Error
}
func (pg *postgres) db() *gorm.DB {
	return pg._db.Debug()
}
func (pg *postgres) AutoMigrate() error {
	db := pg.db()
	tables := []interface{}{
		&models.Movie{},
		&models.User{},
		// May be useful in the future. It is kept here for reference purposes.
		/*
			&database.Token{},
		*/
	}
	err := db.AutoMigrate(tables...)
	return errors.Wrap(err, "db.AutoMigrate")
}
func New(c config.Database) (database.DB, error) {
	db := &postgres{
		conf: c,
	}
	dsn := fmt.Sprintf(DSN, c.Host, c.Port, c.Username, c.Password, c.Name)
	sqlDb, err := gorm.Open(pq.Open(dsn), &gorm.Config{
		//  Doc: GORM perform write (create/update/delete) operations run inside a transaction to ensure data consistency
		// you can disable it during initialization if it is not required, you will gain about 30%+ performance improvement after that
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "db.New")
	}

	db._db = sqlDb

	return db, nil
}
func (pg *postgres) CheckExistingUser(ctx context.Context, email string) (bool, error) {
	var userCount int64
	err := pg.db().Model(&models.User{}).Where(&models.User{Email: email}).Count(&userCount).Error
	return userCount > 0, errors.Wrap(err, "db.CheckExistingUser")
}
func (pg *postgres) CheckUserCredentials(ctx context.Context, email string, password string) (bool, error) {
	var userCount int64
	err := pg.db().Model(&models.User{}).Where(&models.User{Email: email, Password: password}).Count(&userCount).Error
	return userCount > 0, errors.Wrap(err, "db.CheckUserCredentials")
}
func (pg *postgres) CreateUser(ctx context.Context, name string, email string, password string) error {
	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	err := pg.db().Model(&models.User{}).Create(&user).Error
	return errors.Wrap(err, "db.CreateUser")
}
func (pg *postgres) UpdateMovieRating(ctx context.Context, movieName string, rating int8) (float32, error) {
	result := int64(0)
	movie := models.Movie{Name: movieName}
	err := pg.db().Model(&models.Movie{}).Where(&movie).Count(&result).Error
	if err != nil {
		return 0, errors.Wrap(err, "db.UpdateMovieRating")
	}

	if result == 0 {
		movie.Name = movieName
		movie.Rating = float64(rating)
		movie.Count = 1
		err = pg.db().Model(&models.Movie{}).Save(&movie).Error
		if err != nil {
			return 0, errors.Wrap(err, "db.UpdateMovieRating")
		}
		return float32(rating), nil
	}
	err = pg.db().Where(&models.Movie{Name: movieName}).First(&movie).Error
	if err != nil {
		return 0, errors.Wrap(err, "db.UpdateMovieRating")
	}
	curRating := (movie.Rating*float64(movie.Count) + float64(rating)) / float64(movie.Count+1)
	movie.Rating = curRating
	movie.Count += 1

	pg.db().Model(&models.Movie{}).Where(&models.Movie{Name: movie.Name}).Updates(models.Movie{Rating: movie.Rating, Count: movie.Count})
	return float32(curRating), errors.Wrap(err, "db.UpdateMovieRating")
}
func (pg *postgres) GetMoviesData(ctx context.Context) ([]string, error) {
	movies := []models.Movie{}
	err := pg.db().Where(&models.Movie{}).Find(&movies).Error
	res := make([]string, 0, len(movies))
	for _, u := range movies {
		res = append(res, u.Name)
	}
	return res, errors.Wrap(err, "db.GetMoviesData")
}
func (pg *postgres) GetMovieRatings(ctx context.Context) ([]responses.MovieRating, error) {
	movies := []models.Movie{}
	err := pg.db().Where(&models.Movie{}).Find(&movies).Error

	res := make([]responses.MovieRating, 0, len(movies))
	for _, u := range movies {
		res = append(res, responses.MovieRating{Name: u.Name, Rating: u.Rating})
	}
	return res, errors.Wrap(err, "db.GetMovieRatings")
}
