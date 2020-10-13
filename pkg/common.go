package pkg

import "math/rand"

var (
	Seed int
)

// defaults the seed of our Rand Number Generator to 1
func init() {
	InitialiseRandomNumberGenerator(1)
}

// you can manually change the seed of the generator.
// useful for testing purposes / when the goal is to achieve
// the same result on reruns
func InitialiseRandomNumberGenerator(s int) {
	Seed = s
	rand.Seed(int64(Seed))
}

func GetRandomInteger(limit int) int {
	return rand.Intn(limit) //nolint:gosec
}


