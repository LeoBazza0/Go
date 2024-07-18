package main
import "fmt"
func main() {
	/*
fare lo scan di un numero e successivamente vedere se fosse divisibile per 3 o per 5
se il numero avesse il resto della divisione per tre uguale a zero allora si scrive Fizz
se il numero avesse il resto della divisione per cinque uguale a zero allora si scrive Buzz
 	*/

	var num int
	fmt.Print("numero? ")
	fmt.Scan(&num)

	if num%3 == 0 {
		fmt.Println("Fizz")
	}
	if num%5 == 0 {
		fmt.Println("Buzz")
	}
}
