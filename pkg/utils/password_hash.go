package utils

import "golang.org/x/crypto/bcrypt"

func HashPass(pass string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(hash)
}

func ComparePass(hash, pass []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}
