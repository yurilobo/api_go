package main

import "fmt"

type Car struct {
	Name  string
	Model string
}

func main() {
	carro := Car{"Fusca", "VW"}
	fmt.Println(carro.Name)
	carro.Model = "Fiat"
	fmt.Println(carro.Model)
}
