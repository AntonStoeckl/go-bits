package factorial

func PositiveFactorial(number uint8) uint {
	if number == 0 {
		return 1
	}

	result := uint(1)

	for n := uint8(1); n <= number; n++ {
		result *= uint(n)
	}

	return result
}
