package main

import (
	"fmt"

	math "github.com/guialveess/go-utils/packaging/3"
)

func main() {
	m := math.Math{A: 5, B: 3}
	fmt.Println(m.Add())
}
