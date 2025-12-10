package main

import (
	"flag"
	"fmt"
	"shaia.com/project_euler/problems"
)

func main() {
	problem := flag.Int("problem", 0, "The problem number to solve (e.g., 40)")
	flag.Parse()

	if *problem == 0 {
		fmt.Print("Please enter a problem number: ")
		_, err := fmt.Scan(problem)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			return
		}
		if *problem == 0 {
			fmt.Println("No valid problem number provided.")
			return
		}

	}

	solver, exists := problems.Solvers[*problem]
	if !exists {
		fmt.Printf("Problem %d not implemented yet\n", *problem)
		return
	}

	ans := solver()
	fmt.Println(ans)
}
