package main

import "fmt"

func main() {
	var n, count int
	x := 0
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		x += n
		count++
	}
	fmt.Println("la media è: ", float64(x)/float64(count))
}
