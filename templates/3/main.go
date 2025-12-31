package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := tmp.Execute(w, Cursos{
			{Nome: "Go", CargaHoraria: 40},
			{Nome: "Java", CargaHoraria: 60},
			{Nome: "Python", CargaHoraria: 80},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8282", nil)
}
