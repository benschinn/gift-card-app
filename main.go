package main

import (
    "os"
    "fmt"
    "html/template"
    "net/http"
    "net/url"
    "net/http/httputil"
    "bytes"
    "io/ioutil"
)

type PageData struct {
  Title string
  SiteId string
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

var data = PageData{
  Title: os.Getenv("NAME"),
  SiteId: os.Getenv("SITEID"),
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
  <div data-site-id="{{.SiteId}}" data-platform="Other" class="gift-up-target"></div>
  <script type="text/javascript"> (function (g, i, f, t, u, p, s) { g[u] = g[u] || function() { (g[u].q = g[u].q || []).push(arguments) }; p = i.createElement(f); p.async = 1; p.src = t; s = i.getElementsByTagName(f)[0]; s.parentNode.insertBefore(p, s); })(window, document, "script", "https://cdn.giftup.app/dist/gift-up.js", "giftup"); </script>
  `
  t,err := template.New("checkout").Parse(startTmpl + giftUpTmpl + endTmpl)
  if err != nil {
    fmt.Println("error parsing html template")
  }
  t.Execute(w, data)
}

func redeemGiftCard(code string, payload string) *http.Response{
  client := &http.Client{}
  requestURL := url.URL{
    Scheme: "https",
    Host:   "api.giftup.app",
    Path:   "/gift-cards" + code + "/redeem",
  }

  requestHeaders := http.Header{
    "Accept":          {"*/*"},
    "Content-Type":    {"application/json"},
    "Accept-Language": {"en-US,en;q=0.9"},
    "Authorization":   {"bearer " + os.Getenv("API-KEY")},
  }

  jsonBody := []byte(payload)

  request := http.Request{
    Method:        "POST",
    URL:           &requestURL,
    Header:        requestHeaders,
    Body:          ioutil.NopCloser(bytes.NewReader(jsonBody)),
    ContentLength: int64(len(jsonBody)),
  }

  dump, err := httputil.DumpRequest(&request, true)
  if err != nil {
    fmt.Println("dump err", err.Error())
  }

  fmt.Println("******** REQUEST ********")
  fmt.Println(string(dump))

  resp, err := client.Do(&request)
  return resp
}

func main() {
  port := "8090"
  if os.Getenv("PORT") != "" {
    port = os.Getenv("PORT")
  }
  http.HandleFunc("/", landingPage)
  http.HandleFunc("/redeem", redeemPage)
  http.HandleFunc("/giftcards", giftcardPage)
  http.ListenAndServe(":" + port, nil)
}
