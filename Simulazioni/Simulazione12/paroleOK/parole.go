package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var stringMix, stringOk string
	var count int
	var sliceOkFinale, sliceToShow []string
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringMix += scanner.Text() + " "
	}
	sliceMix := strings.Split(stringMix, "")
	for j := 0; j < len(sliceMix); j++ {
		sliceRune := []rune(sliceMix[j])
		if unicode.IsLetter(sliceRune[0]) {
			stringOk += sliceMix[j]
		} else {
			sliceOkFinale = append(sliceOkFinale, stringOk)
			stringOk = ""
		}
	}
	sliceOkFinale = append(sliceOkFinale, stringOk)
	for i := 0; i < len(sliceOkFinale); i++ {
		if sliceOkFinale[i] != "" {
			sliceToShow = append(sliceToShow, sliceOkFinale[i])
			count++
		}
	}
	fmt.Print(sliceToShow)
	fmt.Print(count)

}

//"\n"
