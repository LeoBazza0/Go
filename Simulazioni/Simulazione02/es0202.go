package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	substrings := createSubString(os.Args[1])
	printMap(substrings)
}

func checkFirstAndLast(parola string) bool {
	return parola[0] == parola[len(parola)-1]
}

func createSubString(s string) (m map[string]int) {
	m = make(map[string]int)
	for i := 0; i < len(s); i++ {
		for j := i + 2; j < len(s); j++ {
			substring := s[i : j+1]
			if checkFirstAndLast(substring) {
				m[substring]++
			}
		}
	}
	return
}

func printMap(m map[string]int) {
	keys := getSortedKeys(m)
	for _, key := range keys {
		fmt.Printf("%s -> Occorrenze: %d\n", key, m[key])
	}
}

func getSortedKeys(m map[string]int) (slice []string) {
	for k, _ := range m {
		slice = append(slice, k)
	}
	sort.Slice(slice, func(i, j int) bool {
		return len(slice[i]) > len(slice[j])
	})
	return
}
