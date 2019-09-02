package main

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func imrpmir(valor int) {
	fmt.Println("VALOR ", valor)
}

func main() {
	resultado := somar(3, 4)
	imrpmir(resultado)
}
