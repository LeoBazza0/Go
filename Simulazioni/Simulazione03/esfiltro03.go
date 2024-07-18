/*
package main

import (

	"fmt"
	"os"
	"strconv"

)

	func main() {
		var A, B, C, D int
		for i := 1; i < len(os.Args); i++ {
			num, _ := strconv.Atoi(os.Args[i])
			if num < 18 || num > 30 {
				continue
			} else if num >= 18 && num <= 21 {
				D++
			} else if num >= 22 && num <= 24 {
				C++
			} else if num >= 25 && num <= 27 {
				B++
			} else {
				A++
			}
		}
		Ap := int(float64(A*100) / float64(len(os.Args)))
		Bp := int(float64(B*100) / float64(len(os.Args)))
		Cp := int(float64(C*100) / float64(len(os.Args)))
		Dp := int(float64(D*100) / float64(len(os.Args)))
		fmt.Printf("A : %d %%, \nB : %d %%, \nC : %d %%, \nD : %d %%, \n ", Ap, Bp, Cp, Dp)
	}
*/
package main

import (
	"fmt"
)

const INTERVALS = 4

func main() {
	var counters [INTERVALS]int
	var num, total int
	for {
		fmt.Scanf("%d", &num)
		if num == 0 {
			break
		}
		if num < 18 || num > 30 {
			continue
		}
		counters[rangeOf(num)]++
		total++
	}
	printResults(counters, total)
}

func printResults(counters [INTERVALS]int, total int) {
	//fmt.Println(counters, total)
	for i, c := INTERVALS-1, 'A'; i >= 0; i, c = i-1, c+1 {
		fmt.Printf("%c : %d %%\n", c, int((float64(counters[i])/float64(total))*100))
	}
}

func rangeOf(n int) int {
	return (n - 19) / 3
}
