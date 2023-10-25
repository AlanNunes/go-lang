package main

import (
	"fmt"
	"hello/math"
)

func main() {
	var a int
	var b int
	fmt.Scanf("%d %d", &a, &b)

	resultado := math.Soma(a, b)
	fmt.Printf("Resultado: %v", resultado)
}
