package problems

import "math"

func init() {
	Solvers[40] = Solve40
}

/*
An irrational decimal fraction is created by concatenating the positive integers:
0.12345678910⁢𝟏⁢112131415161718192021⋯

It can be seen that the 12th digit of the fractional part is 1.

If 𝑑𝑛 represents the 𝑛th digit of the fractional part, find the value of the following expression.
𝑑1×𝑑10×𝑑100×𝑑1000×𝑑10000×𝑑100000×𝑑1000000
*/
func Solve40() int64 {
	digits := make(map[int]int)
	lastIndex := 1
	for i := 1; i < 1000000; i++ {
		digits_in_number := getDigitsMSDFirstMath(int64(i))
		for d := range digits_in_number {
			digits[lastIndex] = digits_in_number[d]
			lastIndex++
		}
	}
	return int64(digits[1] * digits[10] * digits[100] * digits[1000] * digits[10000] * digits[100000] * digits[1000000])
}

func getDigitsMSDFirstMath(n int64) []int {
	if n < 0 {
		n = -n // Handle negatives; adjust if you need to preserve sign
	}
	if n == 0 {
		return []int{0}
	}

	// Calculate number of digits
	digitsCount := int(math.Floor(math.Log10(float64(n))) + 1)

	var digits []int
	// Start with the highest power of 10
	pow := int64(math.Pow(10, float64(digitsCount-1)))

	remainder := n
	for pow >= 1 {
		digit := int(remainder / pow)
		digits = append(digits, digit)
		remainder -= int64(digit) * pow
		pow /= 10
	}
	return digits
}
