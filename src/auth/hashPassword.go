package auth

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

const (
	salt string = "lolkek" // некая соль что бы запутать злоумышленика при получение доступа к бд
)

func HashPassword(login, password string) (string, error) {
	password = login + password + salt
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(login, password, hash string) bool {
	password = login + password + salt
	byteHash := []byte(hash)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
