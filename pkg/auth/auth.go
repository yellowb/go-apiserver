package auth

import "golang.org/x/crypto/bcrypt"

// encrypt plain text into hashed text
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// compare hashed text to another plain text, check if they are the same
func Compare(hashedPwd, plainPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
}
