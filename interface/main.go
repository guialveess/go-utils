package main

import "fmt"

type Carro interface {
	ligar(s string) string
}

type MeuCarro struct{}

func (m MeuCarro) ligar(s string) string {
	return s
}

func main() {
	var c Carro = MeuCarro{}
	fmt.Println(
		c.ligar("Carro ligado"),
	)
}
