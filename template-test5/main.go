package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type Profile struct {
	Name	string
	Age		int
}

func (p Profile) ToString() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func main() {
	text := `
{{- .ToString }}
`
	tpl, err := template.New("").Parse(text)
	if err != nil {
		log.Fatal(err)
	}

	p := &Profile{"Tanaka", 31}

	if err := tpl.Execute(os.Stdout, p); err != nil {
		log.Fatal(err)
	}
}