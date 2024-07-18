package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	password := leggipw()
	num, contaLettere, contaMinu, contaMaiu, contaNum := conta(password)
	if num < 6 || num > 15 || contaMinu < 2 || contaMaiu > 3 || contaNum < 2 || num-(contaNum+contaLettere) < 2 {
		fmt.Println("La pw non è definita correttamente:")
	} else {
		fmt.Println("La pw è ben definita!")
	}
	if num < 6 || num > 15 {
		fmt.Println("- La password deve avere lunghezza minima di 6 caratteri e una lunghezza massima 15 caratteri")
	}
	if contaMinu < 2 {
		fmt.Println("- almeno due caratteri della pw devono rappresentare delle lettere miuscole")
	}
	if contaMaiu > 3 {
		fmt.Println("- al massimo 3 caratteri della pw possono rappresentare delle letere maiuscole")
	}
	if contaNum < 2 {
		fmt.Println("- almeno due caratteri della pw devono rappresentare delle cifre decimali")
	}
	if num-(contaNum+contaLettere) < 2 {
		fmt.Println("- almeno due caratteri della pw no devono rappresentare lettere o cifre decimali")
	}
}

func leggipw() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func conta(password string) (num int, contaLettere int, contaMinu int, contaMaiu int, contaNum int) {
	for _, k := range password {
		num++
		if unicode.IsLetter(k) == true {
			contaLettere++
		}
		if unicode.IsLower(k) {
			contaMinu++
		}
		if unicode.IsUpper(k) {
			contaMaiu++
		}
		if unicode.IsNumber(k) == true {
			contaNum++
		}
	}
	return num, contaLettere, contaMinu, contaMaiu, contaNum
}
