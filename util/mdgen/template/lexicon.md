# {{.Owner}}/{{.Name}} lexicon
type: {{.Typ}}


Facts

{{range $fact, $children := .Facts}}
- {{$fact}}
{{end}}
