package main
import "fmt"
func main() {
	/*
	Inserire un voto da uno a cento,
  se non fosse nell'insieme [0;100] esce "voto non valido"
  poi a seconda del numero viene espresso il voto in lettere.
  se maggiore di 90=A; se maggiore di 8=B, se maggiore di 70=C; se maggiore di 60=D; se inferiore di 60=F
	*/

	var voto int
	fmt.Print("voto? ")
	fmt.Scan(&voto)

	if voto < 0 || voto > 100 {
		fmt.Println("voto non valido")
		return //interrompe l'esecuzione del programma
	}

	if voto >= 90 {
		fmt.Println("A")
	} else if voto >= 80 {
		fmt.Println("B")
	} else if voto >= 70 {
		fmt.Println("C")
	} else if voto >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
