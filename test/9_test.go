package test

import (
	"errors"
	"testing"
)

func CheckNegative(value int) error {
	if value < 0 {
		return errors.New("value cannot be negative")
	}
	return nil
}

func TestCheckNegative(t *testing.T) {
	tests := []struct {
		input         int
		expectedError error
	}{
		{input: -1, expectedError: errors.New("value cannot be negative")},
		{input: 1, expectedError: nil},
	}

	for _, test := range tests {
		err := CheckNegative(test.input)
		if !errors.Is(err, test.expectedError) {
			t.Errorf("expected error %v, got %v", test.expectedError, err)
		}
	}
}
