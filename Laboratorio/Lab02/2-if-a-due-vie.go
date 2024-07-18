package main
import "fmt"
func main() {
	/*
	 l'utente scrive un numero e il programma lo scannerizza,
   se nella divizione per 2 non c'è resto, il numero è pari e viene stampato "è pari",
   se nella divizione per 2 c'è resto, il numero è dispari e viene stampato "è dispari"
	*/
	var n int
	fmt.Print("numero: ")
	fmt.Scan(&n)
	if n%2 == 0 {
		fmt.Println(n, "è pari")
	} else {
		fmt.Println(n, "è dispari")
	}
}
