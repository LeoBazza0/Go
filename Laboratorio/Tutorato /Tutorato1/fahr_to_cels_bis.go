package main

import (
	"fmt"
)

func main() {
	var fahr float64
	const FATTORE_SCALA, ZERO = 5.0 / 9.0, 32.0
	for {
		fmt.Scan(&fahr)
		if fahr == 9999 {
			break
		}
		fmt.Println((fahr - ZERO) * FATTORE_SCALA)
	}
}
