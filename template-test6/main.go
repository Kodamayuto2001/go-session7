package main

import (
	"log"
	"os"
	"html/template"
)

func main() {
	text := `
{{- define "contents" }}
<!-- ネスト -->
{{ block "header" . }}
<div>header</div>
{{ end }}
<div>contents</div>
{{ end }}

{{- define "footer" }}
<div>footer</div>
{{ end -}}

{{- block "index" . -}}
<!DOCTYPE html>
<html lang="ja">
<head></head>
<body>
	{{ template "contents" . }}
	{{ template "footer" . }}
</body>
</html>
{{ end -}}
`
	tpl, err := template.New("").Parse(text)
	if err != nil {
		log.Fatal(err)
	}

	if err := tpl.Execute(os.Stdout, nil); err != nil {
		log.Fatal(err)
	}
}