package main
import "fmt"
func main() {
	/*
In questo programma si dichiara una variabile intera,
si  stampa dove si deve inserire il voto, 
si scannerrizza il numero del voto,
se il numero inserito fosse più piccoolo di zero o più grande di trenta, si stampa che il voto non è valido,
altrimenti non si stampa niente
  	*/

	var voto int
	fmt.Print("voto: ")
	fmt.Scan(&voto)
	if voto < 0 || voto > 30 {
		fmt.Println("voto non valido")
	}
}
