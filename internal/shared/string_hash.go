package shared

import "golang.org/x/crypto/bcrypt"

func HashString(s string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	return string(hashBytes), err
}

func CheckHashString(s, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(s), []byte(s))
	return err == nil
}
