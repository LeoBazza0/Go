package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	N, _ := strconv.Atoi(os.Args[1])
	k, _ := strconv.Atoi(os.Args[2])
	//wergb	numIntArr := GeneraNumeri(N, k)
	numIntArrFin := FiltraNumeri(GeneraNumeri(N, k))
	for _, k := range numIntArrFin {
		fmt.Println(k)
	}

}

func GeneraNumeri(N, k int) []int {
	var numIntArr []int
	numStringArr := strconv.Itoa(N)
	for i := 0; i <= len(numStringArr)-k; i++ {
		numString := numStringArr[:i] + numStringArr[i+k:]
		numInt, _ := strconv.Atoi(numString)
		numIntArr = append(numIntArr, numInt)
	}
	return numIntArr
}

func FiltraNumeri(s1 []int) []int {
	var s2 []int
	for i := 0; i < len(s1); i++ {
		if s1[i]%2 != 0 {
			s2 = append(s2, s1[i])
		}
	}
	return s2
}
