package main

import "fmt"

func main() {
	const R, C = 3, 4
	var a [R][C]int

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			a[i][j] *= 2
		}
	}
	fmt.Println(a)
}
