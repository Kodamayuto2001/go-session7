package main

import (
	"log"
	"os"
	"html/template"
)

func main() {
	tpl := template.Must(template.ParseGlob("*.html"))

	if err := tpl.ExecuteTemplate(os.Stdout, "index", nil); err != nil {
		log.Fatal(err)
	}
}