package problems

import (
	"os"
	"strings"
)

// Register the solver automatically.
func init() {
	Solvers[42] = Solve42
}

/*
Solve the coded triangle numbers problem.
The file words.txt contains a list of words.
Count how many words have a name value that is a triangle number.
*/
func Solve42() int64 {
	// Read the words file
	content, err := os.ReadFile("problems\\0042_words.txt")
	if err != nil {
		panic(err) // Handle error appropriately in production
	}
	str := string(content)
	// Assuming the file content is "WORD1","WORD2",...
	words := strings.Split(str, ",")
	for i, w := range words {
		words[i] = strings.Trim(w, "\"")
	}

	// Generate triangle numbers up to i=29 (matching the Python range)
	triangleNums := make(map[int]bool)
	for i := 0; i < 30; i++ {
		t := i * (i + 1) / 2
		triangleNums[t] = true
	}

	// Count words whose value is a triangle number
	count := 0
	for _, word := range words {
		sum := 0
		for _, c := range word {
			if c >= 'A' && c <= 'Z' {
				sum += int(c) - int('A') + 1
			}
		}
		if triangleNums[sum] {
			count++
		}
	}

	return int64(count)
}
