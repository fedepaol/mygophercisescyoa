<!DOCTYPE html>
<html>
  <head>
    <title>.Title</title>
  </head>
  <body>
  <h1>{{ .Title}}</h1>
  {{range .Story}}
        {{ . }}
        <br>
  {{end}}
  {{range .Options}}
        <a href=/{{ .Arc }>{{ .Text }}</a>
        <br>
  {{end}}
  </body>
</html>
