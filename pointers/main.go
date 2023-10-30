package main

import "fmt"

type Calc struct {
	Val int
}

func (calc *Calc) Sum(val int) {
	calc.Val += val
}

func main() {
	calc := Calc{
		Val: 0,
	}
	calc.Sum(1)
	fmt.Print(calc.Val)
}
