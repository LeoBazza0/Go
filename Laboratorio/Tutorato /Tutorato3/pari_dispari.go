package main

import "fmt"

func main() {
	const N = 10
	numeri := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&numeri[i])
	}
	pariDispari(numeri)
}

func pariDispari(numeri []int) {
	for _, k := range numeri {
		if k%2 == 0 {
			fmt.Println(k, "è pari ")
		} else {
			fmt.Println(k, "è dispari")
		}
	}
}
