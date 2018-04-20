package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Slug method for genetate sha hash for user peer connection
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	hasher := md5.New()
	text := time.Now().String()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil)) + string(b)
}
