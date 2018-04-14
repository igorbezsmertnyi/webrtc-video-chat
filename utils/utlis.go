package utils

import (
	"crypto/sha1"
	"encoding/base64"
)

// Slug method for genetate sha hash for user peer connection
func Slug() string {
	hasher := sha1.New()
	hasher.Write(make([]byte, 64))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return sha
}
