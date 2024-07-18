package main

import (
	"fmt"
)

func main() {
	var previous, current int
	fmt.Scan(&current)
	for current != 0 {
		fmt.Scan(&current)
		if current >= previous {
			fmt.Print("+")
		} else {
			fmt.Print("-")
		}
		previous = current
	}
}
