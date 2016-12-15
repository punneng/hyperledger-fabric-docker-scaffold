package main
import (
  "fmt"
  "net/http"
  "io/ioutil"
  "bytes"
  )

func main() {
  url := "http://172.17.0.3:7050/chaincode"
  fmt.Println("URL:>", url)

  var jsonStr = []byte(`{
  "jsonrpc": "2.0",
  "method": "deploy",
  "params": {
    "type": 1,
    "chaincodeID": {
      "path": "https://github.com/IBM-Blockchain/learn-chaincode/finished"
    },
    "ctorMsg": {
      "function": "init",
      "args": [
        "hi there"
      ]
    },
    "secureContext": "test_user30"
  },
  "id": 1
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
