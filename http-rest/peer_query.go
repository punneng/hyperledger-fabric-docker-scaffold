package main
import (
  "fmt"
  "net/http"
  "io/ioutil"
  "bytes"
  )

func main() {
  url := "http://172.17.0.4:7050/chaincode"
  fmt.Println("URL:>", url)

  var jsonStr = []byte(`{
  "jsonrpc": "2.0",
  "method": "query",
  "params": {
    "type": 1,
    "chaincodeID": {
      "name": "3aeb9793d67968f966f2b093c361c70cdbf7a2813a02f7a5da344386580d3b519899b73003b335c587e3d016d44b54eb7d8030bddddbc3e9abf05db81c20eaef"
    },
    "ctorMsg": {
      "function": "read",
      "args": [
        "hello_world"
      ]
    },
    "secureContext": "test_user28"
  },
  "id": 2
}`)

  req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  fmt.Println("response Status:", resp.Status)
  fmt.Println("response Headers:", resp.Header)
  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println("response Body:", string(body))
}
