package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var chiave int
	var testo2 []rune

	fmt.Print("Chiave: ")
	fmt.Scan(&chiave)

	fmt.Print("testo da cifrare: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	testo1 := scanner.Text()

	for _, k := range testo1 {
		if k > 'a' && k < 'z' || k > 'A' && k < 'Z' {
			k += 26
			if k > 'z' && (k-26) <= 'z' {
				k = k - 'a'
			} else if k > 'Z' && (k-26) <= 'Z' {
				k = k - 'A'
			}
			testo2 = append(testo2, k)
		}
		fmt.Println(string(k))
	}
	fmt.Println(testo2)

}
