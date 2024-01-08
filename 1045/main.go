package main

import ("fmt"
	"sort"
)

var pl = fmt.Println

func main() {
	var a float64
	var b float64
	var c float64

	fmt.Scan(&a, &b, &c)
	sides := []float64 {a, b, c}
	sort.Float64s(sides)
	a, b, c = sides[2], sides[1], sides[0]

	if a>= (b + c) {
		pl("NAO FORMA TRIANGULO")
	} else {
		if (a*a) == ((b*b) + (c*c)) {
			pl("TRIANGULO RETANGULO")
		}
		if (a*a) > ((b*b) + (c*c)) {
			pl("TRIANGULO OBTUSANGULO")
		}
		if (a*a) < ((b*b) + (c*c)) {
			pl("TRIANGULO ACUTANGULO")
		}
		if (a == b) && (b == c) {
			pl("TRIANGULO EQUILATERO")
		} else if (a == b) || (b == c) || (a == c) {
			pl("TRIANGULO ISOSCELES")
		}
	}
}