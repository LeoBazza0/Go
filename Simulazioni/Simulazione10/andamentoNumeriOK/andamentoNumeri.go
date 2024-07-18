package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var prec int
	var stringaFinale string
	fmt.Scan(&prec)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		if num < prec {
			stringaFinale += "-"
		} else if num == prec {
			stringaFinale += "="
		} else if num > prec {
			stringaFinale += "+"
		}
		prec = num
	}
	fmt.Println(stringaFinale)
}
