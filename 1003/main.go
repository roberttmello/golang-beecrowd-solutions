package main

import "fmt"

func main() {
	var a int
	var b int
	var soma int
	fmt.Scan(&a)
	fmt.Scan(&b)
	soma = a + b
	fmt.Printf("SOMA = %v\n", soma)
}