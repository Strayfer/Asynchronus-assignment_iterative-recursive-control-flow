package main

import "fmt"

func computePattern(n int) int {
	if n <= 1 {
		return 1
	}
	if n%2 == 1 {
		return n + computePattern(n-1)
	}
	return computePattern(n-2) * 2
}

func main() {
	fmt.Println(computePattern(5)) // change to your M
}
