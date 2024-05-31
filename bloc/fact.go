package bloc

func Factorial(number int) int {
	if number == 0 {
		return 1
	}
	return Factorial(number-1) * number
}
