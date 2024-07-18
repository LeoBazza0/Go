package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	n, _ := strconv.Atoi(os.Args[4])
	fmt.Println(os.Args[1])
	fmt.Println(string(sostituisci([]byte(os.Args[1]), []byte(os.Args[2]), []byte(os.Args[3]), n)))
}

func sostituisci(s, old, new []byte, n int) []byte {
	indexOfOld := nthIndexOfOld(s, old, n)
	if indexOfOld == -1 {
		return s
	}
	return []byte(string(s[:indexOfOld]) + strings.Replace(string(s[indexOfOld:]), string(old), string(new), 1))
}

func nthIndexOfOld(s, old []byte, n int) (index int) {
	sString := string(s)
	oldString := string(old)
	index = strings.Index(sString, oldString)
	for i := 1; i < n; i++ {
		relIndex := strings.Index(sString[index+len(old):], oldString)
		if relIndex == -1 {
			return -1
		}
		index += relIndex + len(old)
	}
	return
}
