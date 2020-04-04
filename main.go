package main

import (
    "os"
    "fmt"
    "net/http"
    "html/template"
)

type Page struct {
  Title string
}

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func html(w http.ResponseWriter, r *http.Request) {
  const tpl = `
  <!DOCTYPE html>
  <html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
  </head>
  <body>
    <div data-site-id="64756a71-b216-4da0-b16a-c8becfcf7a22" data-platform="Other" class="gift-up-target"></div>
    <script type="text/javascript"> (function (g, i, f, t, u, p, s) { g[u] = g[u] || function() { (g[u].q = g[u].q || []).push(arguments) }; p = i.createElement(f); p.async = 1; p.src = t; s = i.getElementsByTagName(f)[0]; s.parentNode.insertBefore(p, s); })(window, document, "script", "https://cdn.giftup.app/dist/gift-up.js", "giftup"); </script>

  </body>
  </html>
  `
  t,err := template.New("webpage").Parse(tpl)
  if err != nil {
    fmt.Println("error parsing html template")
  }
  data := Page{
    Title: os.Getenv("NAME"),
  }
  t.Execute(w, data)
}

func main() {

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)
    http.HandleFunc("/html", html)

    http.ListenAndServe(":8090", nil)
}
