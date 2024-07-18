package main

import "fmt"

func main() {
	slice := []int{2, 7, 18, 4, 1, 58, 36, 23, 15}
	divisore := 3
	fmt.Println(slice)
	fmt.Println("pari:", estraiPari(slice))
	fmt.Println("senza multipli di", divisore, ":", rimuoviMultipli(divisore, slice))
}

func estraiPari(in []int) (out []int) {
	for _, v := range in {
		if v%2 == 0 {
			out = append(out, v)
		}
	}
	return
}

func rimuoviMultipli(m int, in []int) (out []int) {
	for _, v := range in {
		if v%m != 0 {
			out = append(out, v)
		}
	}
	return
}
