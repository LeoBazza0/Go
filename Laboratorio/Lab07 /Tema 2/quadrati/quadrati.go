package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sliceStr := os.Args[1:]
	for i, _ := range sliceStr {
		n, err := strconv.Atoi(sliceStr[i])
		if err == nil {
			if isSquare(n) {
				fmt.Println(n, " ")
			}
		}
	}
}

func isSquare(n int) bool {
	for i := 1; i*i <= n; i++ {
		if n == i*i {
			return true
		}
	}
	return false
}

/*
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	l := len(os.Args)
	if l == 1 {
		fmt.Println("nessun dato")
		os.Exit(1)
	}
	for i := 1; i < l; i++ {
		n, _ := strconv.Atoi(os.Args[i])
		if isSquare(n) {
			fmt.Print(n, " ")
		}
	}
	fmt.Println()
}

func isSquare(n int) bool {
	for i := 1; i*i <= n; i++ {
		if n == i*i {
			return true
		}
	}
	return false
}

*/
