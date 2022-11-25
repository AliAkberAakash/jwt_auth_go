package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GetHash(str []byte) string {
	hash, err := bcrypt.GenerateFromPassword(str, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
