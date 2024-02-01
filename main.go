package main

import "fmt"

type Car struct {
	Name  string
	Model string
}

func (c Car) Andar() {
	fmt.Println("O carro", c.Name, "está andando.")
}

func main() {
	carro := Car{"Fusca", "VW"}
	//fmt.Println(carro.Name)
	carro.Model = "Fiat"
	//fmt.Println(carro.Model)
	carro.Andar()
}
