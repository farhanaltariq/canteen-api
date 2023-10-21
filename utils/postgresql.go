package utils

import (
	"fmt"
	"os"

	"a21hc3NpZ25tZW50/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() error {
	// Get credentials from .env file
	dbCredentials := map[string]string{
		"host":     os.Getenv("DB_HOST"),
		"port":     os.Getenv("DB_PORT"),
		"username": os.Getenv("DB_USERNAME"),
		"password": os.Getenv("DB_PASSWORD"),
		"database": os.Getenv("DB_DATABASENAME"),
	}

	// Construct the DATABASE_URL string
	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		dbCredentials["username"],
		dbCredentials["password"],
		dbCredentials["host"],
		dbCredentials["port"],
		dbCredentials["database"],
	)

	if dbURL == "postgresql://:@:/" {
		dbURL = "postgresql://postgres:aumnQy1gJJHajYnzeXSe@containers-us-west-12.railway.app:5621/railway"
	}

	// connect using gorm pgx
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        dbURL,
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	conn.AutoMigrate(&model.User{}, &model.Canteen{}, &model.Menu{}, &model.CanteenMenu{})
	SetupDBConnection(conn)

	return nil
}

func SetupDBConnection(DB *gorm.DB) {
	db = DB
}

func GetDBConnection() *gorm.DB {
	return db
}
