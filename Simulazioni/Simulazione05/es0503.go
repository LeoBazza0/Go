package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	str := strings.ReplaceAll(line, " ", "")
	count := strings.Count(str, string(str[0]))
	for _, k := range str {
		count1 := strings.Count(str, string(k))
		if count1 != count {
			fmt.Println("non è isogramma")
			os.Exit(0)
		}
	}
	fmt.Println("è isogramma")
}
