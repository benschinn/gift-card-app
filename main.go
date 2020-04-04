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

const startTmpl = `
<!DOCTYPE html>
  <html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
  </head>
  <body>`

const endTmpl= `</body></html>`

var data = Page{
  Title: os.Getenv("NAME"),
}

func landingPage(w http.ResponseWriter, r *http.Request) {
  const tmpl = `<div>landing</div> `
  t,err := template.New("landing").Parse(startTmpl + tmpl + endTmpl)
  if err != nil {
    fmt.Println("error parsing html template")
  }
  t.Execute(w, data)
}

func redeemPage(w http.ResponseWriter, r *http.Request) {
  const tmpl = `<div>redeem</div> `
  t,err := template.New("landing").Parse(startTmpl + tmpl + endTmpl)
  if err != nil {
    fmt.Println("error parsing html template")
  }
  t.Execute(w, data)
}

func giftcardPage(w http.ResponseWriter, r *http.Request) {
  const giftUpTmpl = `
  <div data-site-id="64756a71-b216-4da0-b16a-c8becfcf7a22" data-platform="Other" class="gift-up-target"></div>
  <script type="text/javascript"> (function (g, i, f, t, u, p, s) { g[u] = g[u] || function() { (g[u].q = g[u].q || []).push(arguments) }; p = i.createElement(f); p.async = 1; p.src = t; s = i.getElementsByTagName(f)[0]; s.parentNode.insertBefore(p, s); })(window, document, "script", "https://cdn.giftup.app/dist/gift-up.js", "giftup"); </script>
  `
  t,err := template.New("checkout").Parse(startTmpl + giftUpTmpl + endTmpl)
  if err != nil {
    fmt.Println("error parsing html template")
  }
  t.Execute(w, data)
}

func main() {
    http.HandleFunc("/", landingPage)
    http.HandleFunc("/redeem", redeemPage)
    http.HandleFunc("/giftcards", giftcardPage)
    http.ListenAndServe(":8090", nil)
}
