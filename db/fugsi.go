package db

import (
	"crypto/rand"
	"encoding/hex"
)
// random id
func RandomString() string {
	b := make([]byte, 5)
	rand.Read(b)
	return hex.EncodeToString(b)
}