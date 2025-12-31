package main

func switchType() {
	a := 100
	switch a {
	case 1:
		println("A")
	case 2:
		println("B")
	case 3:
		println("C")
	default:
		println("Outro valor")
	}
}

func main() {
	a := 1
	b := 2
	c := 3

	if a > b || c > a {
		println("A é maior que B ou C é maior que A")
	} else {
		println("Nenhuma das condições é verdadeira")
	}
	switchType()
}
