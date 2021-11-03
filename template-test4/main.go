package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	text := `
{{- createElement "div" "hello world"}}
`
	funcmap := template.FuncMap{
		"createElement": func(tagname, text string) string {
			return fmt.Sprintf("<%s>%s</%s>",tagname,text,tagname)
		},
	}

	tpl := template.New("")

	tpl = tpl.Funcs(funcmap)

	var err error
	tpl, err = tpl.Parse(text)
	if err != nil {
		log.Fatal(err)
	}

	if err := tpl.Execute(os.Stdout, nil); err != nil {
		log.Fatal(err)
	}
}