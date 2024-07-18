package main 

import "fmt"

// mi dice se ci sono lettere latine maiuscole
func hasUpper(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			return true
		}
	}
	return false
}

// mi dice a che posizione è la prima lettera latina maiuscola
func firstUpper(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			return i
			break
		}
	}
	return -1
}

// mi dice a che posizione è l'ultima lettera latina maiuscola
func lastUpper(s string) int {
	for i := len(s)-1; i >= 0; i-- {
		if s[i] >= 'A' && s[i] <= 'Z' {
			return i
			break
		}
	}
	return -1
}

//dice quante cifre, quante lettere o quanti caratteri diversi dai primi due sono contenuti nela stringa inserita
func countDigitsLettersOthers(s string) (int, int, int) {
	var cifre, lettere, altro int
	for _, r := range s {
		if r >= '0' && r <= '9' {
			cifre++
		} else if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			lettere++
		} else {
			altro++
		}
	}
	return cifre, lettere, altro
}

//verifica se una parola è palindroma
func isPalindrome(s string) bool {
    s2 := ""
    for i := len(s)-1; i >= 0; i-- {
        s2 += string(s[i])
    }
    for i := range(s) {
        if s[i] != s2[i] {
            return false
        }
    }
    return true
}

func main() {
	var s string
  fmt.Print ("Inserisci parola: ")
	fmt.Scan(&s)
  if hasUpper(s) {
    fmt.Println ("Ha maiuscole")
  } else {
    fmt.Println("Non ha maiuscole")
  }
  fmt.Println ("La prima maiuscola sta nella posizione", firstUpper(s))
  fmt.Println ("L'ultima maiuscola sta nella posizione", lastUpper(s))
  cifre, lettere, altro := countDigitsLettersOthers(s)
  fmt.Println ("Ci sono", cifre,"cifre, \t", lettere, "lettere, \t", altro, "altri caratteri")
  if isPalindrome(s) {
    fmt.Println ("è palindroma")
  } else {
    fmt.Println("Non è palindroma")
  }
}
