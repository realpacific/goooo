package services

func Factorial(number int) int {
	if number == 1 {
		return 1
	}
	return Factorial(number-1) * number
}
