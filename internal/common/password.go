package common

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	hashedpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashedpass), err

}

func CheckPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
