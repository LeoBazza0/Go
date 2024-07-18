package main

import "fmt"

func main() {
	const N = 10
	numeri := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&numeri[i])
	}
	fmt.Println(minDispari(numeri))
}

func minDispari(numeri []int) int {
	var min int = 99999
	for _, k := range numeri {

		if k%2 == 1 {
			if k < min {
				min = k
			}
		}
	}
	return min
}
