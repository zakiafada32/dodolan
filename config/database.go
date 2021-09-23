package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgreSQL() *gorm.DB {
	dns := "host=localhost user=postgres password=postgres dbname=retail port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
