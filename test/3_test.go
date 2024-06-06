package test

import "testing"

func SumSlice(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func TestSumSlice(t *testing.T) {
	// тест для слайса из положительных чисел
	nums := []int{1, 2, 3, 4, 5}
	expected := 15
	result := SumSlice(nums)
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}

	// тест для пустого слайса
	nums = []int{}
	expected = 0
	result = SumSlice(nums)
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}

	// тест для слайса из отрицательных чисел
	nums = []int{-1, -2, -3, -4, -5}
	expected = -15
	result = SumSlice(nums)
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}

	// тест для слайса из положительных и отрицательных чисел
	nums = []int{-1, 2, -3, 4, -5}
	expected = -3
	result = SumSlice(nums)
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
