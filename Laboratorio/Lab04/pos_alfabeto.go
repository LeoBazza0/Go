package main
import "fmt"
func main ()  {
var s string
fmt.Print ("Inserisci una sequenza di simboli con un '.' finale: ")
fmt.Scan (&s)
for i:=0; i<len(s); i++ {
  if s[i]==46 {
    break
  } else if s[i]>=97 && s[i]<=122 {
    fmt.Println ( string(s[i]), "-", "è la", s[i]-96, "^a")
    } else if s[i]>=48 && s[i]<=57 {
      fmt.Println (string(s[i]), "-", string(s[i]))
    }else {
      fmt.Println (string(s[i]), " - altro")
    }
  }
}
