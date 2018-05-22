# {{.Owner}}/{{.Name}} {{.Typ}} lexicon

##  {{.Name}} facts
{{$name := .Name}}
{{$owner := .Owner}}
{{range $fact, $children := .Facts}}<details><summary>{{$name}}.{{$fact}}</summary><p>

#### Example of finding every {{$fact}} and having a review bot comment on it:

```{{$name}}
tenets:
  - name: find_all_{{$fact}}
    doc:  Example query to find all instances of {{$fact}}
    bots:
      codelingo/review
        comment: This is a {{$fact}}.
    query: |
      import {{$owner}}/{{$name}}

      @ review.comment
      {{$name}}.{{$fact}}
```
</p></details>
{{end}}
