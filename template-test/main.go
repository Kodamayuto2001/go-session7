package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	text := "{{ . }}\n"

	//	テキストをパースしてExecute関数の第二引数で値を渡す
	//	渡した値はテキスト内の{{}}で「.」として参照することができる。
	tpl, err := template.New("").Parse(text)
	if err != nil {
		log.Fatal(err)
	}

	value := "hello world"

	if err := tpl.Execute(os.Stdout, value); err != nil {
		log.Fatal(err)
	}
}