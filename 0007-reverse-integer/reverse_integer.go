package reverse_integer

func reverse(x int) int {
	var (
		intMax = 2147483647
		intMin = -2147483648
		y      = 0
	)
	if x < intMin || x > intMax {
		return 0
	}
	for x != 0 {
		y = y*10 + x%10
		x /= 10
		if x > 0 && y > intMax/10 {
			return 0
		} else if x < 0 && y < intMin/10 {
			return 0
		}
	}
	return y
}
