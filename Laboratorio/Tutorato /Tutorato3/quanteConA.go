package main

import "fmt"

func main() {
	const N = 10
	parole := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&parole[i])
	}
	fmt.Println(quanteConA(parole))
}

func quanteConA(parole []string) int {
	var count int = 0
	for _, k := range parole {
		for _, j := range k {
			if j == 'a' {
				count++
			}
		}
	}
	return count
}
