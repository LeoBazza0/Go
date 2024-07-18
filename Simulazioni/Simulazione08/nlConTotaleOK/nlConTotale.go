package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var count int
	file, _ := os.Open(os.Args[1])
	scanner1 := bufio.NewScanner(file)
	for scanner1.Scan() {
		count++
	}
	i := 1
	file2, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner2 := bufio.NewScanner(file2)
	for scanner2.Scan() {
		fmt.Printf("%d/%d %s\n", i, count, scanner2.Text())
		i++
	}
}
