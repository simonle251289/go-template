package utils

import "golang.org/x/crypto/bcrypt"

func HashString(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		return password
	}
	return string(bytes)
}

func ValidateHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
