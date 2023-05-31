package utils

import (
	"math/rand"
	"time"
)


// ----- Random ------

var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

// Random int in range [0, n)
func RandInt(n int) int {
	return seed.Intn(n)
}


// ------ Prime ------

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}
