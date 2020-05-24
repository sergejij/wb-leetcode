package main

func myAtoi(str string) (res int) {
	number, i, minus, strLen := 0, 0, 1, len(str)
	for ; i < strLen && str[i] == ' '; i++ {}
	if i < strLen && str[i] == '-' {
		minus = -1
	}
	if i < strLen && (str[i] == '-' || str[i] == '+') {
		i++
	}
	for ; i < strLen && (str[i] >= 48 && str[i] <= 57); i++ {
		number = (number * 10) + int(str[i]) - 48
		if minus == -1 && number/2147483648 != 0 {
			return -2147483648
		}
		if minus == 1 && number/2147483647 != 0 {
			return 2147483647
		}
	}
	return number * minus
}
