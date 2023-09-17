package util

import (
	"errors"
	coreError "mobile-ecommerce/internal/core/error"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("error when encrypt password: " + err.Error())
	}

	return string(encryptedPassword), nil
}

// return nil on success, error when failure
func CheckPassword(encryptedPassword string, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password)); err != nil {
		logrus.Error("error when compare password:", err)
		return coreError.ErrWrongPassword
	}
	return nil
}
