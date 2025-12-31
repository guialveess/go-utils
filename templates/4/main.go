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
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	tmp := template.Must(template.New("content.html").ParseFiles(templates...))
	err := tmp.Execute(os.Stdout, Cursos{
		{Nome: "Go", CargaHoraria: 40},
		{Nome: "Java", CargaHoraria: 60},
		{Nome: "Python", CargaHoraria: 80},
	})
	if err != nil {
		panic(err)
	}
}
