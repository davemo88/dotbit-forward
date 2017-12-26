package main

import (
  "fmt"
  "net/http"
//  "strings"
  "bytes"
  "io/ioutil"
  "encoding/json"
)

const NAMECOIND_HOST = "http://127.0.0.1:8336"

type nameValue map[string]string

type rpcResponse struct {
  Result        []nameValue `json:"result"`
  Id            string `json:"id"`
  Error         string `json:"error"`
}

func DotBitForward(w http.ResponseWriter, r *http.Request) {
//  domainParts := strings.Split(r.Host, ".")
  domainParts := [2]string{"hg", "taco"}
  jsonStr := fmt.Sprintf(`{"jsonrpc":"1.0","id":"gotext","method":"name_filter","params":["^d/%v$"]}`, domainParts[0])
  req, err := http.NewRequest("POST", NAMECOIND_HOST, bytes.NewBuffer([]byte(jsonStr)))
  req.Header.Add("content-type","text/plain")
  req.SetBasicAuth("rpcuser","tacos")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }
  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Fprintln(w, string(body))
  var data rpcResponse
  _ = json.Unmarshal(body, &data)
  fmt.Fprintln(w, data.Result[0]["value"])
//  http.Redirect(w,r,"http://google.com",301)
}

func main() {
    http.HandleFunc("/", DotBitForward)
    http.ListenAndServe(":8080", nil)
}
