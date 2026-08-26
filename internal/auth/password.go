package auth

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(v string) string {
	h := sha256.Sum256([]byte("sanzoujin:" + v))
	return hex.EncodeToString(h[:])
}
func CheckPassword(hash, plain string) bool { return hash == HashPassword(plain) }
