package main

import (
	"net/http"
	"text/template"
)

type Curso struct	{
	Nome   				string
	CargaHoraria 	int
}

type Cursos []Curso


func main(){

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		t:= template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, Cursos{
			{"Go", 70},
			{"Java", 50},
			{"Python", 90},
		})
		if err != nil {
		panic(err)
		}
	})
	http.ListenAndServe(":8282", nil)
	
}