package hasher

import (
	"server/UserService/app/apperror"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns the hash password before going in the db.
func HashPassword(input string) (string, error) {
	if len(input) == 0 {
		return "", apperror.ErrorEmptyField
	}

	post_password := []byte(input)

	hashed, err := bcrypt.GenerateFromPassword(post_password, 12)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

// ComparePassword compares the password with the hash password.
func ComparePassword(input string, hashed string) (bool, error) {
	if len(input) == 0 {
		return false, apperror.ErrorEmptyField
	}

	post_password := []byte(input)
	hashed_password := []byte(hashed)

	err := bcrypt.CompareHashAndPassword(hashed_password, post_password)

	return err == nil, err
}
