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

func hasKey(m map[string]string, s string) bool {
  _, ok := m[s]
  return ok
}

func getIpFromMap(mapVal string, subdomain string) string {
  return "bogus"
}

func getDotBitDomain(requestHost string) (string, string) {
  return "", "hg"
}

//func getRedirectIp() {
//  var redirectIp string
//
//  switch {
//  case hasKey(dotBitRecord, "ip"):
//    redirectIp = dotBitRecord["ip"]
//  case hasKey(dotBitRecord, "map"):
//    redirectIp = getIpFromMap(dotBitRecord["map"],dotBitSubdomain)
//  case hasKey(dotBitRecord, "ns"):
//  }
//
//  ftm.Fprintln(w, redirectIp)
//}

func dotBitForward(w http.ResponseWriter, r *http.Request) {
  dotBitSubdomain, dotBitDomain := getDotBitDomain(r.Host)
  jsonStr := fmt.Sprintf(`{"jsonrpc":"1.0","id":"gotext","method":"name_filter","params":["^d/%v$"]}`, dotBitDomain)
  req, err := http.NewRequest("POST", NAMECOIND_HOST, bytes.NewBuffer([]byte(jsonStr)))
  req.Header.Add("content-type","text/plain")
  req.SetBasicAuth("rpcuser","tacos")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }
  body, _ := ioutil.ReadAll(resp.Body)
//  fmt.Fprintln(w, string(body))
  var data rpcResponse
  _ = json.Unmarshal(body, &data)
//  fmt.Fprintln(w, data.Result[0]["value"])

  var dotBitRecord map[string]string
  _ = json.Unmarshal([]byte(data.Result[0]["value"]), &dotBitRecord)

  var redirectIp string

  switch {
  case hasKey(dotBitRecord, "ip"):
    redirectIp = dotBitRecord["ip"]
  case hasKey(dotBitRecord, "map"):
    redirectIp = getIpFromMap(dotBitRecord["map"],dotBitSubdomain)
  case hasKey(dotBitRecord, "ns"):
  }

  fmt.Fprintln(w, redirectIp)

}

func main() {
    http.HandleFunc("/", dotBitForward)
    http.ListenAndServe(":8080", nil)
}
