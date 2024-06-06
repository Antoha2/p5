package test

import (
	"errors"
	"testing"
)

func CheckSign(value int) (string, error) {
	if value < 0 {
		return "", errors.New("value is negative")
	} else if value == 0 {
		return "zero", nil
	} else {
		return "positive", nil
	}
}

func TestCheckSign(t *testing.T) {
	tests := []struct {
		input          int
		expectedResult string
		expectedError  error
	}{
		{input: 5, expectedResult: "positive", expectedError: nil},
		{input: 0, expectedResult: "zero", expectedError: nil},
		{input: -3, expectedResult: "", expectedError: errors.New("value is negative")},
	}

	for _, test := range tests {
		result, err := CheckSign(test.input)

		if result != test.expectedResult {
			t.Errorf("for input %d, expected result %v, got %v", test.input, test.expectedResult, result)
		}

		if !errors.Is(err, test.expectedError) {
			t.Errorf("for input %d, expected error %v, got %v", test.input, test.expectedError, err)
		}
	}
}
