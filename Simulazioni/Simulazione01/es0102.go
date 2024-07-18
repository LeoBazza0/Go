package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	check(s)
	substrings := Sottostringhe(s)
	fmt.Println("output:")
	printMap(substrings)

}

func check(s string) {
	_, err := strconv.Atoi(s)
	if err != nil {
		os.Exit(0)
	}

}

func Sottostringhe(s string) map[string]int {
	substrings := make(map[string]int)
	for j, _ := range s {
		for i := j + 1; i < len(s); i++ {
			substring := s[j : i+1]
			if isTheOrderCorrect(substring) {
				substrings[substring]++
			}
		}
	}
	return substrings
}

func isTheOrderCorrect(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}

func printMap(m map[string]int) {
	keys := sortKeys(m)
	for _, key := range keys {
		fmt.Printf("%s %d\n", key, m[key])
	}
}

func sortKeys(m map[string]int) (keys []string) {
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return len(keys[i]) > len(keys[j])
	})
	return
}
