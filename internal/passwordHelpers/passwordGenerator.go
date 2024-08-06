package passwordHelpers

import (
	"golang.org/x/crypto/argon2"
)

func hashPass(salt []byte, plainPassword string) []byte {
	hashedPass := argon2.IDKey([]byte(plainPassword), salt, 1, 64*1024, 4, 32)
	return append(salt, hashedPass...)

}

func generateRandomBytes(n int) []byte {
	return make([]byte, n)
}

func PasswordGenerator(plainPassword string) []byte {
	salt := generateRandomBytes(8)
	return hashPass(salt, plainPassword)
}
