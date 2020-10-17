package pkg

import (
	"math/rand"
	"time"
)

var (
	Seed int64
)

// defaults the seed of our Rand Number Generator to 1
func init() {
	InitialiseRandomNumberGenerator(time.Now().UTC().UnixNano())
}

// you can manually change the seed of the generator.
// useful for testing purposes / when the goal is to achieve
// the same result on reruns
func InitialiseRandomNumberGenerator(s int64) {
	Seed = s
	rand.Seed(Seed)
}

func GetRandomInteger(limit int) int {
	if limit == 0 {
		return 0
	}
	return rand.Intn(limit) //nolint:gosec
}

func GetRandomIntegerInRange(lower, upper int) int {
	return GetRandomInteger(upper-lower)+lower
}

