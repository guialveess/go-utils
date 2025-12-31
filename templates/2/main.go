package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	tmp := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := tmp.Execute(os.Stdout, Cursos{
		{Nome: "Go", CargaHoraria: 40},
		{Nome: "Java", CargaHoraria: 60},
		{Nome: "Python", CargaHoraria: 80},
	})
	if err != nil {
		panic(err)
	}
}
