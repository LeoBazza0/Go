package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	slice := strings.Split(scanner.Text(), " ")
	for i := 0; i < len(slice); i++ {
		int, _ := strconv.Atoi(slice[i+1])
		for j := 0; j < int; j++ {
			fmt.Print(slice[i])
		}
		i++
	}

}
