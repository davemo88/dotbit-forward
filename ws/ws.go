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

type stringMap map[string]string

type rpcResponse struct {
  Result        []stringMap `json:"result"`
  Id            string `json:"id"`
  Error         string `json:"error"`
}

func dotBitForward(w http.ResponseWriter, r *http.Request) {
  dotBitSubdomain, dotBitDomain := getDotBitDomain(r.Host)
  req := getRpcRequest(dotBitDomain)
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }
  dotBitRecord := getDotBitRecord(resp)
  redirectIp := getRedirectIp(dotBitRecord, dotBitSubdomain)
  http.Redirect(w, r, redirectIp, 303)
}

func getDotBitDomain(requestHost string) (string, string) {
  return "", "2ez"
}

func getRpcRequest(dotBitDomain string) *http.Request {
  jsonStr := fmt.Sprintf(`{"jsonrpc":"1.0","id":"gotext","method":"name_filter","params":["^d/%v$"]}`, dotBitDomain)
  req, _ := http.NewRequest("POST", NAMECOIND_HOST, bytes.NewBuffer([]byte(jsonStr)))
  req.Header.Add("content-type","text/plain")
  req.SetBasicAuth("rpcuser","tacos")
  return req
}

func getDotBitRecord(resp *http.Response) stringMap{
  body, _ := ioutil.ReadAll(resp.Body)
//  fmt.Fprintln(w, string(body))
  var data rpcResponse
  _ = json.Unmarshal(body, &data)
//  fmt.Fprintln(w, data.Result[0]["value"])

  var dotBitRecord stringMap
// why is Result an array
  _ = json.Unmarshal([]byte(data.Result[0]["value"]), &dotBitRecord)

  return dotBitRecord
}

func getRedirectIp(dotBitRecord stringMap, dotBitSubdomain string) string {
  var redirectIp string

  switch {
  case dotBitSubdomain == "" && hasKey(dotBitRecord, "ip"):
    redirectIp = dotBitRecord["ip"]
  case dotBitSubdomain == "" && hasKey(dotBitRecord, "map"):
    redirectIp = getIpFromMap(dotBitRecord["map"], dotBitSubdomain)
  case hasKey(dotBitRecord, "ns"):
//    nameservers := dotBitRecord["ns"]
  case hasKey(dotBitRecord, "translate"):
  }

  return redirectIp
}

func getIpFromMap(mapVal string, subdomain string) string {
  var m stringMap
  _ = json.Unmarshal([]byte(mapVal), &m)
  var ip string
  if hasKey(m, subdomain) {
    ip = m[subdomain]
  } else {
    ip = ""
  }
  return ip
}

func hasKey(m stringMap, s string) bool {
  _, ok := m[s]
  return ok
}


func main() {
    http.HandleFunc("/", dotBitForward)
    http.ListenAndServe(":8080", nil)
}
