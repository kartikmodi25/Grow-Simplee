package postgres

import (
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

	// db._db = sqlDb

	// return db, nil

	// db, err := sql.Open("postgres")

	// if err != nil {
	// 	return nil, errors.Wrap(err, "db.GetConnection")
	// }
	return sqlDb, nil
}
