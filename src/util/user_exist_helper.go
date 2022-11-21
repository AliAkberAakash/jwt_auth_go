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
