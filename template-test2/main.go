package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	text := `Name: {{ .Name }}
Age: {{ .Age }}
`
	tpl, err := template.New("").Parse(text)
	if err != nil {
		log.Fatal(err)
	}

	//	field名は大文字で始める
	v := struct {
		Name	string
		Age		int
	}{
		Name:	"Tanaka",
		Age:	32,
	}

	if err := tpl.Execute(os.Stdout, v); err != nil {
		log.Fatal(err)
	}
}