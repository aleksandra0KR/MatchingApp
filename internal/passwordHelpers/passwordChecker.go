package passwordHelpers

import "bytes"

func CheckPass(passHash string, plainPassword string) bool {
	passHashBytes := []byte(passHash)
	salt := passHashBytes[:8]
	userPassHash := hashPass(salt, plainPassword)
	return bytes.Equal(userPassHash, passHashBytes)
}
