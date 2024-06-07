package test

import (
	"errors"
	"testing"
)

func CheckSign(value int) (string, error) {
	if value < 0 {
		return "", errors.New("<0")
	} else if value == 0 {
		return "0", nil
	} else {
		return ">0", nil
	}
}

func TestCheckSign(t *testing.T) {
	tests := []struct {
		input          int
		expectedResult string
		expectedError  error
	}{
		{input: 5, expectedResult: "<0", expectedError: nil},
		{input: 0, expectedResult: "0", expectedError: nil},
		{input: -3, expectedResult: "", expectedError: errors.New(">0")},
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
