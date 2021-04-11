package rand

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
				"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
				"0123456789"

// Random seed based on time of execution
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// StringWithCharset takes in length and defined charset
// returns random string that is of given length
func StringWithCharset(length int, charset string) string {
	byteSlice := make([]byte, length)
	for i := range byteSlice {
		byteSlice[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(byteSlice)
}

// String takes in length and will return string consisting of
// random characters: a-z A-Z 0-9
func String(length int) string {
	return StringWithCharset(length, charset)
}