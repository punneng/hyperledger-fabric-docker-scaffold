package main
import (
  "fmt"
  "net/http"
  "io/ioutil"
  "bytes"
  )

func main() {
  url := "http://172.17.0.4:7050/registrar"
  fmt.Println("URL:>", url)

  var jsonStr = []byte(`{
    "enrollId": "test_user28",
    "enrollSecret": "hackathon128"
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
