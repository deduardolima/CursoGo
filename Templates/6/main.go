package main

import (
	"html/template"
	"os"
	"strings"
)

type Curso struct	{
	Nome   				string
	CargaHoraria 	int
}

type Cursos []Curso


func main(){
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t:= template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
	t = template.Must(t.ParseFiles(templates...))

	err := t.Execute(os.Stdout, Cursos{
			{"Go", 70},
			{"Java", 50},
			{"Python", 90},
		})
	if err != nil {
		panic(err)
		}
	
}

