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

const NMCD_LINKNAME = "NMCD"

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
  domainParts := strings.Split(requestHost, ".")
  var dotBitSubdomain, dotBitDomain string
  switch {
    case len(domainParts) == 2:
      dotBitSubdomain = ""
      dotBitDomain = domainParts[0]
    case len(domainParts) == 3:
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

func getDotBitRecord(resp *http.Response) stringMap{
  body, _ := ioutil.ReadAll(resp.Body)
  var data rpcResponse
  _ = json.Unmarshal(body, &data)

  var dotBitRecord stringMap
// why is Result an array
  if len(data.Result) > 0 {
  _ = json.Unmarshal([]byte(data.Result[0]["value"]), &dotBitRecord)
  }
  fmt.Println(data)

  return dotBitRecord
}

func getRedirectIp(dotBitRecord stringMap, dotBitSubdomain string) string {
  var redirectIp string

  switch {
  case dotBitSubdomain == "" && hasKey(dotBitRecord, "ip"):
    redirectIp = dotBitRecord["ip"]
  case dotBitSubdomain != "" && hasKey(dotBitRecord, "map"):
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
  switch {
  case hasKey(m, subdomain):
    ip = m[subdomain]
  case hasKey(m, "*"):
    ip = m["*"]
  default:
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
