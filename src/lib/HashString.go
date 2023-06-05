package lib

import "golang.org/x/crypto/bcrypt"

func HashString(some_string string) (string, error) {
	some_bytes := []byte(some_string)
	hashed_bytes, err := bcrypt.GenerateFromPassword(some_bytes, 8)
	if err != nil {
		return "", err
	}
	hashed_string := string(hashed_bytes)
	return hashed_string, nil
}