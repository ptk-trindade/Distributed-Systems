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

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	
	for i := 2; i*i < n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}