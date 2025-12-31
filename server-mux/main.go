package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", &blog{title: "My Blog"})
	http.ListenAndServe(":8080", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", HelloGuilherme)
	http.ListenAndServe(":9090", mux2)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func HelloGuilherme(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Guilherme!"))
}

type blog struct {
	title string
}

func (b *blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Blog"))
}
