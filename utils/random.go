package utils

import (
	"fmt"
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // min + a random numbe between 0 -> max-min hence eventually returning a number which is in between min and max
}

//generates a random string comprising n chars
func RandomString(n int) string {
	var s string
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(len(alphabet))]
		s += string(c)
	}

	return s
}

// RandomOwner returns a random set of 6 chars
func RadnomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return randomInt(0, 1000)
}

//RadomCurrency returns generates a random curreny
func RadomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

//Generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(6))
}
