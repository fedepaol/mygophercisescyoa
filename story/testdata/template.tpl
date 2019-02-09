<!DOCTYPE html>
<html>
  <head>
    <title>.Node.Title</title>
  </head>
  <body>
  <h1>{{ .Node.Title}}</h1>
  {{range .Node.Story}}
  {{ . }}{{end}}
  {{range .Node.Options}} <a href={{ $.URL }}/{{ .Arc }}>{{ .Text }}</a>
  {{end}}</body>
</html>
