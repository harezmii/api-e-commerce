package hash

import "golang.org/x/crypto/bcrypt"

func PasswordHash(password string) (string, error) {
	hashPassword, hashPasswordError := bcrypt.GenerateFromPassword([]byte(password), 13)
	if hashPasswordError != nil {
		return "", hashPasswordError
	}
	return string(hashPassword), nil
}

func PasswordHashCompare(password, hashedPassword string) bool {
	passwordHashCompareError := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if passwordHashCompareError != nil {
		return false
	}
	return true
}
