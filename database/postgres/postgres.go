package postgres

import (
	"backend-assignment/database/models"
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
	fmt.Println(dsn, sqlDb)
	if err != nil {
		return nil, errors.Wrap(err, "db.New")
	}
	return sqlDb, nil
}
func CheckExistingUser(db *gorm.DB, email string) (bool, error) {
	var userCount int64
	err := db.Model(&models.User{}).Where(&models.User{Email: email}).Count(&userCount).Error
	return userCount > 0, err
}
func CheckUserCredentials(db *gorm.DB, email string, password string) (bool, error) {
	var userCount int64
	err := db.Model(&models.User{}).Where(&models.User{Email: email, Password: password}).Count(&userCount).Error
	return userCount > 0, err
}
func GenerateUserToken(db *gorm.DB, email string, accessToken string) error {
	UserToken := models.UserToken{
		Email:       email,
		AccessToken: accessToken,
	}
	err := db.Model(&models.UserToken{}).Create(&UserToken).Error
	return err
}
func CreateUser(db *gorm.DB, name string, email string, password string) error {
	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	err := db.Model(&models.User{}).Create(&user).Error
	return err
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
