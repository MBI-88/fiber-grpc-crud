package tools

import "golang.org/x/crypto/bcrypt"

func GenerateHasKey(key string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(key),10)
	return string(hash), err
}

func CheckHahPassword(userhash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userhash), []byte(password))
	return err
}