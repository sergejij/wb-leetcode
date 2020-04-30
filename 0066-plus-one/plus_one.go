package main

func plusOne(digits []int) []int {
	arr, lenDigits := make([]int, 1), len(digits)
	for i := lenDigits - 1; i >= 0; i-- {
		if digits[i] == 9 && i == 0 {
			arr[0], digits[i] = 1, 0
			arr = append(arr, digits...)
			return arr
		} else if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1
			return digits
		}
	}
	return digits
}
