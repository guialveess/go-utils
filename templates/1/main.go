package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Golang", 40}
	tmp := template.Must(template.New("Curso Template").Parse("Curso: {{.Nome}}, Carga Horária: {{.CargaHoraria}} horas")) // cria o template
	err := tmp.Execute(os.Stdout, curso)                                                                                   // escreve no console
	if err != nil {
		panic(err)
	}
}
