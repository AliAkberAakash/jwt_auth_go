package util

import (
	"fmt"
	"jwt-auth/main/src/data"
	"jwt-auth/main/src/dto"
)

func GetUserFromDB(email string) (dto.User, error) {

	var foundUser dto.User

	for _, u := range data.Users {
		if u.Email == email {
			foundUser = u
			return foundUser, nil
		}
	}

	return foundUser, fmt.Errorf("User not Found")
}
