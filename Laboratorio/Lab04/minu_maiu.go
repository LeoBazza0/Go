package main
import (
	"fmt"
	"unicode"
)

func main() {
	var parola     string
	var	minu, maiu bool
  fmt.Print("inserisci parola: ")
	fmt.Scan(&parola)
	for _, c := range parola {
		if unicode.IsLower(c) {
			minu = true
		} else if unicode.IsUpper(c) {
			maiu = true
		}
		if maiu && minu {
			break
		}
	}
	fmt.Println("minu:", minu, "maiu:", maiu)
}
