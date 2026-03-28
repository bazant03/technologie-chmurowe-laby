package main

import (
    "html/template"
    "net/http"
    "os"
)

type PageData struct {
    IP       string
    Hostname string
    Version  string
}

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
    hostname, _ := os.Hostname()
    data := PageData{
        IP:       "127.0.0.1",
        Hostname: hostname,
        Version:  VVV,
    }
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    if err := tpl.Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    http.HandleFunc("/", handler)
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    http.ListenAndServe(":"+port, nil)
}