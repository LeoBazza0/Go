package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var max, min, count, tot int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		count++
		num := scanner.Text()
		numInt, _ := strconv.Atoi(num)
		if count == 1 {
			max = numInt
			min = numInt
		} else if numInt > max {
			max = numInt
		} else if numInt < min {
			min = numInt
		}
		tot += numInt
	}
	fmt.Println("min:", min)
	fmt.Println("max:", max)
	fmt.Println("media:", float64(float64(tot)/float64(count)))

}
