package main

import "fmt"

func main() {
	result, err := soma(1, 2)
	fmt.Println(result, err)
}
func soma(a int, b int) (int, error) {
	if a+b > 10 {
		return 0, fmt.Errorf("soma maior que 10")
	}
	return a + b, nil
}
