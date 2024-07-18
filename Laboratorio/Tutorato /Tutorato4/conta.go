package main

import (
	"fmt"
	"os"
)

func main() {

	for _, k := range os.Args[1:] {
		for i := 0; i < len(k)-1; i++ {
			if k[i] == k[i+1] {
				fmt.Println("la prima parola con una doppia è: ", k)
				os.Exit(0)
			}
		}
	}
}
