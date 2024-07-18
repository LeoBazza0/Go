package main
import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strings"
)

type Cerchio struct {
nome string
x,y, raggio float64
}

func main () {
var descr string
var cPrec Cerchio
scanner := bufio. NewScanner (os.Stdin)
for scanner. Scan () {
descr=scanner.Text ()
c1:= newcircle (descr)
fmt. Print (c1)
//distPunti:= distanzaPunti (cPrec.x, cPrec.y, c1.x, c1.y )
//fmt. Printin (distPunti)
if tocca (cPrec, c1) {
fmt. Print(" tangente. " )
} else {
fmt.Print (" non tangente,")
}
if maggiore (cPrec, c1) {
fmt.Printin (" maggiore")
} else {
fmt. Printin (" minore o uguale")
}
cPrec=c1


}
}

func newcircle (descr string) Cerchio {
slice1:= strings.Split (descr, " ")
var c Cerchio
c.nome= slice1[0]
c.X, _= strconv. ParseFloat (slice1[1], 64)
c.y, _= strconv. ParseFloat (slice1 (2], 64)
c.raggio, _= strconv. ParseFloat (slice1 [3], 64)
return c
}

func distanzaPunti (x1, y1, x2, y2 float64) float64 {
dist:=math.Sqrt((X1-×2)*(×1-×2)+(y1-y2)*(y1-y2))
return dist
}











func tocca (cerc1, cerc2 Cerchio) bool {
var tangenti bool
if distanzaPunti(cerc1.x, cerc1.y, cerc2.x, cerc2.y)>cerci.raggio+cerc2.raggio {
return tangenti
} else if distanzaPunti (cerc1.x, cerc2.x, cerc1.y, cerc2.y)=(cerci.raggio+cerc.raggio) {
tangenti=true
return tangenti
} else {
diff:=math.Abs(cerc1.raggio-cerc2.raggio)
if distanzaPunti (cerc1.x, cerc2.x, cerc1.y, cerc2.y)=diff {
bool=true
return tangenti
} else {
bool=false
return bool
}
}
}










func maggiore (cerc1, cerc2 Cerchio) bool {
pi:=3.141592
area1:=pi*cerc1.raggio* cerc1.raggio
area2:= pi*cerc2.raggio*cerc2. raggio
if area1<area2{
return true
} else {
return false
}
}
