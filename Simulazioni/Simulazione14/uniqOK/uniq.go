package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var current string
	count := 0
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == current {
			count++
		} else {
			if count > 0 {
				fmt.Printf("%d %s\n", count+1, current)
			} else if current != "" {
				fmt.Println(count+1, current)
			}
			current = line
			count = 0
		}
	}
	if count > 0 {
		fmt.Printf("%d %s\n", count+1, current)
	} else if current != "" {
		fmt.Println(current)
	}

}
