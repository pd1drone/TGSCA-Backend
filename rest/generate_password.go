package rest

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

func GeneratePassword() (string, error) {
	const length = 8
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b)[:length], nil
}

func MD5HashPassword(pass string) string {

	// Create an MD5 hash object
	hash := md5.New()

	// Write the input string to the hash object
	hash.Write([]byte(pass))

	// Get the final hash sum as a byte slice
	hashSum := hash.Sum(nil)

	// Convert the hash sum to a hexadecimal string
	hashString := hex.EncodeToString(hashSum)

	return hashString

}
