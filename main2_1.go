package main

import "fmt"

func main() {
	var N, temp, sum, ideal, over int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&temp)
		sum += temp
		if temp >= 195 && temp <= 210 {
			ideal++
		} else if temp > 210 {
			over++
		}
	}
	fmt.Println(ideal, over, float64(sum)/float64(N))
}
