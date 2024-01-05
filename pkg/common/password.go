package common

import "golang.org/x/crypto/bcrypt"

func PasswordHash(plainText string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(h), nil
}

func CheckPassword(plainText string, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plainText))
}
