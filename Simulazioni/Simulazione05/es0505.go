package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num1, _ := strconv.Atoi(os.Args[1])
	num2, _ := strconv.Atoi(os.Args[2])
	fmt.Print(num1 * num2 / mcd(num1, num2))
}

// minimo comune divisore con Euclide
func mcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
