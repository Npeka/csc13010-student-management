package crypto

import "golang.org/x/crypto/bcrypt"

func GetHash(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(hashed)
}
