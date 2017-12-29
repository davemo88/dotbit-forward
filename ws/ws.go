package main

import (
  "fmt"
  "net/http"
  "strings"
  "bytes"
  "io/ioutil"
  "encoding/json"
  "os"
)

type stringMap map[string]string

type rpcResponse struct {
  Result        []stringMap `json:"result"`
  Id            string `json:"id"`
  Error         string `json:"error"`
}

func dotBitForward(w http.ResponseWriter, r *http.Request) {
  dotBitSubdomain, dotBitDomain := getDotBitDomain(r.Host)
  fmt.Printf("%s => %s, %s\n", r.Host, dotBitSubdomain, dotBitDomain)
  req := getRpcRequest(dotBitDomain)
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }
  data, dotBitRecord := getDotBitRecord(resp)
  redirectIp := getRedirectIp(dotBitRecord, dotBitSubdomain)
  if len(redirectIp) > 0 {
    redirectDest := fmt.Sprintf("http://%s", redirectIp)
    fmt.Println("redirecting to", redirectDest)
    http.Redirect(w, r, redirectDest, 303)
  } else {
    fmt.Fprintln(w, "no dice, got", data.Result)
  }

}

func getDotBitDomain(requestHost string) (string, string) {
  domainParts := strings.Split(requestHost, ".")
  var dotBitSubdomain, dotBitDomain string
  switch {
    case len(domainParts) == 3:
      dotBitSubdomain = ""
      dotBitDomain = domainParts[0]
    case len(domainParts) == 4:
      dotBitSubdomain = domainParts[0]
      dotBitDomain = domainParts[1]
    default:
      dotBitSubdomain = ""
      dotBitDomain = "2ez"
  }
  return dotBitSubdomain, dotBitDomain
}

func getRpcRequest(dotBitDomain string) *http.Request {
  jsonStr := fmt.Sprintf(`{"jsonrpc":"1.0","id":"gotext","method":"name_filter","params":["^d/%v$"]}`, dotBitDomain)
  req, _ := http.NewRequest("POST", os.Getenv("NMCD_HOST"), bytes.NewBuffer([]byte(jsonStr)))
  req.Header.Add("content-type","text/plain")
  req.SetBasicAuth(os.Getenv("RPCUSER"),os.Getenv("RPCPASSWORD"))
  return req
}

func getDotBitRecord(resp *http.Response) (rpcResponse, stringMap) {
  fmt.Println(resp)
  body, _ := ioutil.ReadAll(resp.Body)
  var data rpcResponse
  _ = json.Unmarshal(body, &data)

  var dotBitRecord stringMap
// why is Result an array
  if len(data.Result) > 0 {
  _ = json.Unmarshal([]byte(data.Result[0]["value"]), &dotBitRecord)
  }

  return data, dotBitRecord
}

func getRedirectIp(dotBitRecord stringMap, dotBitSubdomain string) string {
  var redirectIp string

  switch {
  case dotBitSubdomain == "" && hasKey(dotBitRecord, "ip"):
    redirectIp = dotBitRecord["ip"]
  case hasKey(dotBitRecord, "map"):
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
  fmt.Println(m, subdomain)
  switch {
  case hasKey(m, subdomain):
    ip = m[subdomain]
  case hasKey(m, "*"):
    ip = m["*"]
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
