package main

import "fmt"

// Punto è un punto nel piano cartesiano
type Punto struct {
	x, y float64
}

func main() {
	var p1, p2 Punto
  var x, y float64

	fmt.Print("x e y di P1: ")
	fmt.Scan(&p1.x, &p1.y)

	fmt.Print("x e y di P2: ")
	fmt.Scan(&p2.x, &p2.y)

	fmt.Println(funzUno(p1, p2))

	funzDue(&p1, 2, 3)
	fmt.Println(p1)
  fmt.Print ("inserisci due valori x e y: ")
  fmt.Scan (&x, &y)
fmt.Println (newPunto(x, y))
  ppp:= specchiaPunto (&p1, p1)
  fmt.Println (ppp)
}

func funzUno(p1, p2 Punto) Punto {
	var p Punto
	p.x = (p1.x + p2.x) / 2
	p.y = (p1.y + p2.y) / 2
	return p
}

func funzDue(p *Punto, horiz, vert float64) {
	(*p).x += horiz
	(*p).y += vert
}

func newPunto(x, y float64) (p Punto) {
p.x = x
p.y = y
return
}

func specchiaPunto(p *Punto, pp Punto) {
(*p).x = - pp.x
}
