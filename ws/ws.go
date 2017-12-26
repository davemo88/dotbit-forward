package main

import (
  "fmt"
  "net/http"
  "strings"
  "bytes"
  "io/ioutil"
)

const NAMECOIND_HOST = "http://0.0.0.0:8336"

type Subdomains map[string]http.Handler
func (subdomains Subdomains) ServeHTTP(w http.ResponseWriter, r *http.Request) {

  domain_parts := strings.Split(r.Host, ".")
  fmt.Fprintf(w, domain_parts[0])
  jsonStr := []byte(`{"jsonrpc":"1.0","id":"gotext","method":"getinfo","params":[]}`)
  req, err := http.NewRequest("POST", NAMECOIND_HOST, bytes.NewBuffer(jsonStr))
  req.Header.Add("content-type","text/plain")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }
  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Fprintf(w,string(body))
}


func main() {
    subdomains := make(Subdomains)
    http.ListenAndServe(":8080", subdomains)
}
