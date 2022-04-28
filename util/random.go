package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const _ = "абвгдеёжзиклмнопрстуфхцчшщъыьэюя"
const alphabetEng = "abcdefghijklmnopqrstuwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomInt generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabetEng)

	for i := 0; i < n; i++ {
		c := alphabetEng[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generates random user name with length 6
func RandomUser() string {
	return RandomString(6)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
