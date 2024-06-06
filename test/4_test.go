package test

import (
	"math/rand"
	"testing"
	"testing/quick"
	"time"
)

func GenerateRandomSlice(length int) []int {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	slice := make([]int, length)
	for i := range slice {
		slice[i] = r.Intn(100)
	}
	return slice
}

func TestGenerateRandomSlice(t *testing.T) {
	f := func(length int) bool {
		if length < 0 {
			return true
		}

		slice := GenerateRandomSlice(length)
		if len(slice) != length {
			return false
		}

		for _, v := range slice {
			if v < 0 || v >= 100 {
				return false
			}
		}
		return true
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
