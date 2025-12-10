package problems

// Register the solver automatically.
func init() {
	Solvers[39] = Solve39
}

/*
Solve
If 𝑝 is the perimeter of a right angle triangle with integral length sides, {𝑎,𝑏,𝑐},
there are exactly three solutions for 𝑝 =120.
{20,48,52}, {24,45,51}, {30,40,50}
For which value of 𝑝 ≤1000, is the number of solutions maximised?
*/
func Solve39() int64 {
	solution := make(map[int]int)
	maxVal, maxP := 0, 0
	for p := 1; p <= 1000; p++ {
		for a := 1; a <= p/3; a++ {
			for b := a; b <= (p-a)/2; b++ {
				c := p - a - b
				if isTriangle(a, b, c) && isRightAngle(a, b, c) {
					solution[p]++
				}
			}
		}
	}

	for p, val := range solution {
		if val > maxVal {
			maxP, maxVal = p, val
		}
	}
	return int64(maxP)
}

func isTriangle(a, b, c int) bool {
	if a+b > c && a+c > b && b+c > a {
		return true
	}
	return false
}

func isRightAngle(a, b, c int) bool {
	return a*a+b*b == c*c
}
