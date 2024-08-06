package passwordHelpers

import "bytes"

func CheckPass(passHash []byte, plainPassword string) bool {
	salt := passHash[:8]
	userPassHash := hashPass(salt, plainPassword)
	return bytes.Equal(userPassHash, passHash)
}
