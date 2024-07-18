package main

import (
	"fmt"
)

func main() {
	var symbol string

	_, err := fmt.Scan(&symbol)
	if err != nil {
		fmt.Println(err)
	}

	switch symbol {
	case "L":
		for i := 1; i <= 4; i++ {
			fmt.Println("*")
		}
		fmt.Println("*****")
	case "T":
		fmt.Println("*****")
		line := "  *"
		for i := 1; i <= 4; i++ {
			fmt.Println(line)
		}
	case "Z":
		fmt.Println("*****")
		line := "   *"
		for i := 1; i <= 3; i++ {
			fmt.Println(line)
			line = line[1:]
		}
		fmt.Println("*****")
	default:
		fmt.Println("input non valido")

	}
}
