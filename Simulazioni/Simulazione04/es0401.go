package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])
	for i := n1; i <= n2; i++ {
		fmt.Print(printD(i), " ")
	}
}

func printD(i int) string {
	if i%105 == 0 {
		return "DinDonDan"
	} else if i%35 == 0 {
		return "DonDan"
	} else if i%21 == 0 {
		return "DinDan"
	} else if i%15 == 0 {
		return "DinDon"
	} else if i%7 == 0 {
		return "Dan"
	} else if i%5 == 0 {
		return "Don"
	} else if i%3 == 0 {
		return "Din"
	} else {
		return strconv.Itoa(i)
	}
}
