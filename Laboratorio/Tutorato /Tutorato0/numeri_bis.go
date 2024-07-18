package main

import "fmt"

func main() {
	var n int
	x := 0
	for i := 0; i < 10; i++ {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		x += n
	}
	fmt.Println(x)
}
