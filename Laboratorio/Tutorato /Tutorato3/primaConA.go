package main

import "fmt"

func main() {
	const N = 10
	parole := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&parole[i])
	}
	fmt.Println(primaConA(parole))
}

func primaConA(parole []string) string {
	for _, k := range parole {
		for _, j := range k {
			if j == 'a' {
				return k
			}
		}
	}
	return ""
}
