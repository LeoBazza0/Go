package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var count int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	for _, k := range line {
		if strings.ContainsAny(string(k), "aeiouAEIOU") {
			count++
		}
	}
	fmt.Printf("ci sono %d vocali", count)

}
