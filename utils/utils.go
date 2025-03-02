package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword returns a bcrypt hashed password or error if any
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// IsPasswordValid checks if the given hashed password and the password
// provided by the user matched
func IsPasswordValid(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
