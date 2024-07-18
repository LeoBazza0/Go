package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var prodotto int = 1
	var sliceNumeri []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		numeri := scanner.Text()
		sliceNumeri = strings.Split(numeri, " ")
	}

	for i := 0; ; i++ {
		n, _ := strconv.Atoi(sliceNumeri[i])
		if n == 0 {
			break
		}
		if n != 0 && n%2 == 0 {
			prodotto *= n
		}
	}

	fmt.Println("il prodotto è:", prodotto)
}
