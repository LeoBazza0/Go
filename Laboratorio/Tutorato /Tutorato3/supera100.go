package main

import (
	"fmt"
	"os"
)

func main() {
	var num int
	for {
		fmt.Scan(&num)
		if num == -1 {
			os.Exit(0)
		} else if num >= 100 {
			fmt.Println("il primo numero che supera il 100 è: ", num)
			//os.Exit(0)
			//break
			return
		} else if num < 100 {
			continue
		}
	}
}

//FUNZIONA MA NON BENISSIMO
