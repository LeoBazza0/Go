package main

import (
	"fmt"
	"math"
)

func main() {
	const N = 20
	var target int
	numeri := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&numeri[i])
	}
	fmt.Print("inserrisci il target: ")
	fmt.Scan(&target)

	fmt.Println(numeroVicino(numeri, target))
}

func numeroVicino(numeri []int, target int) int {
	var closest int = numeri[0]
	for i := 0; i < len(numeri); i++ {
		if math.Abs(float64(target-numeri[i])) < math.Abs(float64(target-closest)) {
			closest = numeri[i]
		}
	}
	return closest
}
