package main

import (
	"fmt"
)

func main() {
	var s1, s2 string
	fmt.Println("Inesrisci la prima frase: ")
	fmt.Scan(&s1)

	fmt.Println("Inserisci la seconda frase: ")
	fmt.Scan(&s2)

	if isAnagram(s1, s2) {
		fmt.Println("sono anagrammi")
	} else {
		fmt.Println("non sono anagrammi")
	}

}

func isAnagram(s1, s2 string) bool {
	sm1 := make(map[string]int, len(s1))
	sm2 := make(map[string]int, len(s2))

	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		sm1[string(s1[i])]++
		fmt.Println(sm1)
	}

	for i := 0; i < len(s2); i++ {
		sm1[string(s2[i])]--
		fmt.Println(sm1)
	}

	for i := 0; i < len(s1); i++ {
		if sm2[string(s1[i])] != 0 {
			return false
		}
	}
	return true
}
