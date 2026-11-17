package main

import "fmt"

func main() {
	K := 7 // change to your K
	total := 0
	for i := 1; i <= K; i++ {
		if i%2 == 0 {
			total += i * 2
		} else {
			total -= i
		}
	}
	fmt.Println(total)
}
