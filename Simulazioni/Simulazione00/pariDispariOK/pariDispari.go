package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("nessun valore in ingresso")
		os.Exit(0)
	}
	seque := os.Args[1:]
	//	fmt.Println(seque)
	for i := 0; i < len(seque)-1; i++ {
		numi1, err1 := strconv.Atoi(seque[i])
		numi2, err2 := strconv.Atoi(seque[i+1])
		//	fmt.Println(numi1, numi2)
		if err1 != nil {
			fmt.Printf("elemento in posizione %d non valido\n", i+1)
			return
		} else if err2 != nil {
			fmt.Printf("elemento in posizione %d non valido\n", i+2)
			return
		}
		if err1 != nil || err2 != nil || (numi1+numi2)%2 == 0 {
			fmt.Printf("elemento in posizione %d non valido\n", i+2)
			return

		}
	}
	fmt.Println("sequenza valida")
}
