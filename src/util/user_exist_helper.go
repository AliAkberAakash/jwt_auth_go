package util

import (
	"fmt"
	"jwt-auth/src/dto"
	"log"

	"gorm.io/gorm"
)

func GetUserFromDB(email string, db *gorm.DB) (*dto.User, error) {

	var user *dto.User

	tx := db.Where("email = ?", email).First(&user)

	// checks if there is any error
	if tx.Error != nil {
		log.Print(tx.Error.Error())
		return user, tx.Error
	}

	// checks the count of query
	var count int64
	tx.Count(&count)

	if count == 0 {
		log.Print(count)
		return user, fmt.Errorf("User not found")
	}

	log.Print(count)
	log.Print(user)

	return user, nil
}
