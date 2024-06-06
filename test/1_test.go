package test

import (
	"testing"
)

func add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	// Таблица тестовых случаев
	testCases := []struct {
		a       int
		b       int
		wantSum int
	}{
		{1, 2, 3},
		{5, -3, 2},
		{0, 0, 0},
		{-1, -1, -2},
		{100, 200, 300},
	}

	for _, testCase := range testCases {
		gotSum := add(testCase.a, testCase.b)

		if gotSum != testCase.wantSum {
			t.Errorf("add(%d, %d): got %d, want %d", testCase.a, testCase.b, gotSum, testCase.wantSum)
		}
	}
}
