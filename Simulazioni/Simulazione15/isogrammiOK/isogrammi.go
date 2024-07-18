package main

import (
	"bufio"
	"fmt"
	"os"
)

func isIsogram(s string) bool {
	frequenze := make(map[rune]int)
	for _, c := range s {
		frequenze[c]++
		if c != ' ' && c != '-' && frequenze[c] > 1 {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Print(line, ": ")
		if isIsogram(line) {
			fmt.Println("SI")
		} else {
			fmt.Println("NO")
		}
	}
}
