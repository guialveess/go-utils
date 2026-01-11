package main

import (
	"fmt"

	"github.com/guialveess/go-utils/packaging/1/math"
)

func main() {

	m := math.Math{A: 5, B: 3}

	fmt.Println(m.Add())
	fmt.Println(m.Subtract())
	fmt.Println(m.Multiply())
	fmt.Println(m.Divide())
}
