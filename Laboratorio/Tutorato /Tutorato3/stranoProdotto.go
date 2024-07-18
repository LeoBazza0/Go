package main

import "fmt"

func main() {
	const N = 10
	numeri := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&numeri[i])
	}
	fmt.Println(stranoProdotto(numeri))
}

func stranoProdotto(numeri []int) int {
	var tot int = 1
	for _, k := range numeri {
		if k > 7 && k%4 == 0 {
			tot *= k
		}
	}
	return tot
}
