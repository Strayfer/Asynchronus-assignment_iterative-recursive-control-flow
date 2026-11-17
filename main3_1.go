package main

import "fmt"

func countEvenDigits(n int) int {
	if n == 0 {
		return 0
	}
	if (n%10)%2 == 0 {
		return 1 + countEvenDigits(n/10)
	}
	return countEvenDigits(n / 10)
}

func main() {
	fmt.Println(countEvenDigits(48271))
	fmt.Println(countEvenDigits(1359))
}
