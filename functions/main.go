package main

import "fmt"

func main() {
	//res := somaTudo(1, 2, 3, 4, 5)
	somaTudo := func(val ...int) func() int {
		res := 0
		for _, v := range val {
			res += v
		}
		return func() int {
			return res
		}
	}

	fmt.Printf("Resultado: %d", somaTudo(1, 2)())
}

// func somaTudo(val ...int) (res int) {
// 	for _, v := range val {
// 		res += v
// 	}
// 	return
// }
