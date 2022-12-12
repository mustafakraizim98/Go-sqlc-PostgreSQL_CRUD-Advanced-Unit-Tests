package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomHeadquarters generates a random location for the headquarters
func RandomHeadquarters() string {
	headquarterses := []string{"Jerusalem, Palestine",
		"London, United Kingdom",
		"Washington, United States of America",
		"Ottava, Canada",
		"Ankara, Turkey"}
	n := len(headquarterses)
	return headquarterses[rand.Intn(n)]
}

// RandomFoundationYear generates a random selection of foundation year
func RandomFoundationYear() int64 {
	return RandomInt(2000, 2022)
}
