package main

import "fmt"

func main()  {
	const PI = 3.14159
	var raio float64
	fmt.Scan(&raio)
	var area = PI * (raio * raio)
	fmt.Printf("A=%.4f\n", area)
}