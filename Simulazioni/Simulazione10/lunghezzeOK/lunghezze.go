package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var line string
	var m map[int]string
	m = make(map[int]string)
	if len(os.Args) != 2 {
		fmt.Println("Fornire 1 nome di file")
		os.Exit(0)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("File non accessibile")
		os.Exit(0)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line += scanner.Text() + " "
	}
	sliceparole := strings.Fields(line)
	for i := 0; i < len(sliceparole); i++ {
		m[len(sliceparole[i])] += sliceparole[i] + " "
	}
	printMap(m)
}

func printMap(m map[int]string) {
	maxNumber := max(m)
	for i := 1; i <= maxNumber; i++ {
		if m[i] == "" {
			continue
		}
		fmt.Println(i, ":", m[i])
	}
}
func max(numbers map[int]string) int {
	var maxNumber int
	for n := range numbers {
		maxNumber = n
		break
	}
	for n := range numbers {
		if n > maxNumber {
			maxNumber = n
		}
	}
	return maxNumber
}
