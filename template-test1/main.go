package main

import (
	"log"
	"os"
	"text/template"
)

//	テキスト内にある【.name】と【.age】をそれぞれmapの値で置き換える
func main() {
	text := `Name: {{ .name }}
Age: {{ .age }}
`
	
	tpl, err := template.New("").Parse(text)
	if err != nil {
		log.Fatal(err)
	}

	m := map[string]interface{}{
		"name":	"Tanaka",
		"age":	31,
	}

	if err := tpl.Execute(os.Stdout, m); err != nil {
		log.Fatal(err)
	}
}