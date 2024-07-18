package main

import (
	"fmt"
	"os"
)

func main() {
	var max int = 0
	var e int
	stringa_iniziale := os.Args[1:]
	for i := 0; i < len(stringa_iniziale); i++ {
		if len(stringa_iniziale[i]) > max {
			max = len(stringa_iniziale[i])
			e = i
		} else if len(stringa_iniziale[i]) == max {
			e = i
		}
	}
	fmt.Println(stringa_iniziale[e])
}
