package main

import (
	"fmt"
	"io"
)

func main() {
	var stringa string
	var sliceS []string
	for {
		_, err := fmt.Scan(&stringa)
		if err == io.EOF {
			break
		}
		sliceS = append(sliceS, stringa)

	}
	fmt.Print(frequenze(sliceS))

}

func frequenze(s []string) map[string]int {
	mappa := make(map[string]int)
	for i := 0; i < len(s); i++ {
		mappa[s[i]]++
	}
	return mappa
}
