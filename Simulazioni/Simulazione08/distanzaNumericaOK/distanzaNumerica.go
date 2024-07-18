package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var tot float64
	tot, _ = strconv.ParseFloat(os.Args[1], 64)
	args2, _ := strconv.ParseFloat(os.Args[2], 64)
	tot += args2
	for i := 1; i < len(os.Args)-2; i++ {
		elem1, _ := strconv.ParseFloat(os.Args[i], 64)
		elem2, _ := strconv.ParseFloat(os.Args[i+1], 64)
		elem3, _ := strconv.ParseFloat(os.Args[i+2], 64)
		if math.Abs(elem1-elem2) == math.Abs(elem2-elem3) {
			tot += elem3
		} else {
			fmt.Println("NON OK")
			os.Exit(0)
		}
	}
	fmt.Println(tot)
}
