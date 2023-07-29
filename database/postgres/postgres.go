package postgres

import (
	"backend-assignment/database/models"
	"backend-assignment/responses"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	pq "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgres struct {
	_db *gorm.DB
}

const DSN = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Kolkata"

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Movie{}, &models.User{})
	return errors.Wrap(err, "db.AutoMigrate")
}
func GetConnection() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := fmt.Sprintf(DSN, os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	sqlDb, err := gorm.Open(pq.Open(dsn), &gorm.Config{
		//  Doc: GORM perform write (create/update/delete) operations run inside a transaction to ensure data consistency
		// you can disable it during initialization if it is not required, you will gain about 30%+ performance improvement after that
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "db.GetConnection()")
	}
	return sqlDb, nil
}
func CheckExistingUser(db *gorm.DB, email string) (bool, error) {
	var userCount int64
	err := db.Model(&models.User{}).Where(&models.User{Email: email}).Count(&userCount).Error
	return userCount > 0, errors.Wrap(err, "db.CheckExistingUser")
}
func CheckUserCredentials(db *gorm.DB, email string, password string) (bool, error) {
	var userCount int64
	err := db.Model(&models.User{}).Where(&models.User{Email: email, Password: password}).Count(&userCount).Error
	return userCount > 0, errors.Wrap(err, "db.CheckUserCredentials")
}
func CreateUser(db *gorm.DB, name string, email string, password string) error {
	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	err := db.Model(&models.User{}).Create(&user).Error
	return errors.Wrap(err, "db.CreateUser")
}
func UpdateMovieRating(db *gorm.DB, movieName string, rating int8) (float32, error) {
	result := int64(0)
	movie := models.Movie{Name: movieName}
	err := db.Model(&models.Movie{}).Where(&movie).Count(&result).Error
	if err != nil {
		return 0, errors.Wrap(err, "db.UpdateMovieRating")
	}

	if result == 0 {
		movie.Name = movieName
		movie.Rating = float64(rating)
		movie.Count = 1
		err = db.Model(&models.Movie{}).Save(&movie).Error
		if err != nil {
			return 0, errors.Wrap(err, "db.UpdateMovieRating")
		}
		return float32(rating), nil
	}
	err = db.Where(&models.Movie{Name: movieName}).First(&movie).Error
	if err != nil {
		return 0, errors.Wrap(err, "db.UpdateMovieRating")
	}
	curRating := (movie.Rating*float64(movie.Count) + float64(rating)) / float64(movie.Count+1)
	movie.Rating = curRating
	movie.Count += 1

	db.Model(&models.Movie{}).Where(&models.Movie{Name: movie.Name}).Updates(models.Movie{Rating: movie.Rating, Count: movie.Count})
	return float32(curRating), errors.Wrap(err, "db.UpdateMovieRating")
}
func GetMoviesData(db *gorm.DB) ([]string, error) {
	movies := []models.Movie{}
	err := db.Where(&models.Movie{}).Find(&movies).Error
	res := make([]string, 0, len(movies))
	for _, u := range movies {
		res = append(res, u.Name)
	}
	return res, errors.Wrap(err, "db.GetMoviesData")
}
func GetMovieRatings(db *gorm.DB) ([]responses.MovieRating, error) {
	movies := []models.Movie{}
	err := db.Where(&models.Movie{}).Find(&movies).Error

	res := make([]responses.MovieRating, 0, len(movies))
	for _, u := range movies {
		res = append(res, responses.MovieRating{Name: u.Name, Rating: u.Rating})
	}
	return res, errors.Wrap(err, "db.GetMovieRatings")
}
