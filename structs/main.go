package main

import "fmt"

type Engine struct {
	horsePower int
}

func (engine Engine) Start() {
	fmt.Printf("Starting with %d of horse power", engine.horsePower)
}

type Car struct {
	Engine
	Name string
}

func main() {
	fordKa := Car{
		Name: "Ford Ka",
		Engine: Engine{
			horsePower: 10,
		},
	}

	fordKa.Start()
}
