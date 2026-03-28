package main

import (
    "html/template"
    "net/http"
    "os"
)

// PageData zawiera pola używane w szablonie HTML
type PageData struct {
    IP       string
    Hostname string
    Version  string
}

// prosty HTML template zwracający informacje o serwerze
var tpl = template.Must(template.New("page").Parse(`
<!doctype html>
<html>
  <head><meta charset="utf-8"><title>Info</title></head>
  <body>
    <h1>Info</h1>
    <p><strong>IP:</strong> {{.IP}}</p>
    <p><strong>Hostname:</strong> {{.Hostname}}</p>
        <p><strong>Wersja:</strong> {{.Version}}</p>
  </body>
</html>
`))

func handler(w http.ResponseWriter, r *http.Request) {
    // pobierz nazwę hosta
    hostname, _ := os.Hostname()
    // przygotuj dane do szablonu
    data := PageData{
        IP:       "127.0.0.1",
        Hostname: hostname,
        Version:  VVV,
    }
    // ustaw nagłówek i wyrenderuj szablon
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    if err := tpl.Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    // zarejestruj handler i uruchom serwer (PORT lub 8080)
    http.HandleFunc("/", handler)
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    http.ListenAndServe(":"+port, nil)
}