package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var host = os.Getenv("POSTGRES_HOST")
var user = os.Getenv("POSTGRES_USER")
var password = os.Getenv("POSTGRES_PASSWORD")
var dbname = os.Getenv("POSTGRES_DB")
var port = os.Getenv("POSTGRES_PORT")
var sslMode = os.Getenv("POSTGRES_SSL")

var connectionString = fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
	host,
	user,
	password,
	dbname,
	port,
	sslMode,
)

func CreateConnection() *gorm.DB {
	db, err := gorm.Open(
		postgres.Open(connectionString),
		&gorm.Config{},
	)

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
