package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	for i := 1; i <= N; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("COMPUTATION")
		case i%3 == 0:
			fmt.Println("ALGORITHM")
		case i%5 == 0:
			fmt.Println("LOGIC")
		default:
			fmt.Println(i)
		}
	}
}
