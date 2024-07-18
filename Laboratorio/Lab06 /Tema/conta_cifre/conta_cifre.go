package main
import (
  "fmt"
  "unicode"
  "os"
  "bufio"
  "strings"
)

func main() {
	numStringheConCifre := 0
	var contatoreCifre [10]int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "stop" {
			break
		}
		paroleInserite := strings.Split(line, " ")
		for i := 0; i < len(paroleInserite); i++ {
			if contaCifre(paroleInserite[i], &contatoreCifre) {
				numStringheConCifre++
			}
		}
	}
	fmt.Printf("%d stringhe con cifre\n", numStringheConCifre)
	fmt.Println("[0 1 2 3 4 5 6 7 8 9]")
	fmt.Println(contatoreCifre)
}

func contaCifre(s string, numCifre *[10]int) (haCifre bool) {
	for _, r := range s {
		if unicode.IsNumber(r) {
			haCifre = true
			numCifre[r - '0']++
		}
	}
	return
}
