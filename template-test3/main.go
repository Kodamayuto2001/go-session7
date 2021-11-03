package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	text := "A {{- .B -}} C {{- .D -}} E\n"

	tpl, err := template.New("").Parse(text)
	if err != nil {
		log.Fatal(err)
	}

	m := map[string]interface{}{
		"B":	"B",
		"D":	"D",
	}

	if err := tpl.Execute(os.Stdout, m); err != nil {
		log.Fatal(err)
	}
}