package test

import (
	"math/rand"
	"testing"
	"testing/quick"
)

func GenerateRandomNumber(min, max int) int {
	if min > max {
		panic("min cannot be greater than max")
	}

	return rand.Intn(max-min+1) + min
}

func TestGenerateRandomNumber(t *testing.T) {
	f := func(min, max int) bool {
		n := GenerateRandomNumber(min, max)
		return min <= n && n <= max
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
