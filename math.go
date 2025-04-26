package main

import "fmt"

func main() {
	fmt.Println(Soma(2, 2))
}

func Soma(a int, b int) int {
	return a + b
}

func Subtracao(a int, b int) int {
	return a - b
}

func Multiplicacao(a int, b int) int {
	return a * b
}
