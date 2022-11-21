package data

import (
	"fmt"
	"jwt-auth/src/dto"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitStore() (*gorm.DB, error) {

	pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
	)

	log.Print(pgConnString)

	DB, err := gorm.Open(postgres.Open(pgConnString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&dto.User{})
	if err != nil {
		log.Fatal(err)
	}

	return DB, nil
}
