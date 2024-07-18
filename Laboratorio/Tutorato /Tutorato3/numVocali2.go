package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var max, numI int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fullline := scanner.Text()
	line := strings.Split(fullline, " ")
	for i := 0; i < len(line); i++ {
		count := numVocali(line[i])
		if count > max {
			max = count
			numI = i
		}
	}
	fmt.Print("la parola con più vocali è: ", line[numI])
}

func numVocali(parola string) (count int) {
	for _, k := range parola {
		if strings.ContainsAny(string(k), "aeiouAEIOU") {
			count++
		}
	}
	return count

}
