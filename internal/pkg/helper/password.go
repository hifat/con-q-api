package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (hash string, err error) {
	newHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(newHash), err
}
