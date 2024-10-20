package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const connectionString = "host=localhost user=postgres password=hui dbname=postgres port=32768 sslmode=disable TimeZone=Asia/Shanghai"

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
