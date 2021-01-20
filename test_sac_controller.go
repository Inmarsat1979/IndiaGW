package main

import (
  "fmt"
  "strings"
  "time"
  "net/http"
  "io/ioutil"
)

func send_outside_curl_command () {
  url := "http://localhost:8080/set-terminal-data"
  method := "POST"

  payload_str := `
  {
        "TIMEMS" : 1606632300000,
        "DID" : 335614417,
        "LAT": 13.480099678039551,
        "LON": 77.144599914550781,
        "SAC": 122,
        "SVNS":[
        {
                "SVN" : 1014,
                "SAT0IP":  "100.81.7.236",
                "SAT0SUBNET":  "255.255.255.255",
                "ETH0IP" :  "10.61.34.38",
                "ETH0SUBNET": "255.255.255.252"
        },
        {
                "SVN" : 1013,
                "SAT0IP":  "100.81.7.236",
                "SAT0SUBNET":  "255.255.255.255",
                "ETH0IP" :  "10.61.50.38",
                "ETH0SUBNET": "255.255.255.252"
        },
        {
                "SVN" : 1011,
                "SAT0IP":  "100.72.3.28",
                "SAT0SUBNET":  "255.255.255.255",
                "ETH0IP" :  "192.168.1.6",
                "ETH0SUBNET": "255.255.255.0"
        }
        ]
  }`
  payload := strings.NewReader(payload_str)
  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)
  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("DID", "335577991")
  req.Header.Add("SAC", "193")
  req.Header.Add("Authorization", "Basic YWRtaW46YWRtaW4=")
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))



}


func send_inside_curl_command () {


  url := "http://localhost:8080/set-terminal-data"
  method := "POST"

  payload_str := `
  {
        "TIMEMS" : 1606632300000,
        "DID" : 335614417,
        "LAT": 13.480099678039551,
        "LON": 77.144599914550781,
        "SAC": 121,
        "SVNS":[
        {
                "SVN" : 1014,
                "SAT0IP":  "100.81.7.236",
                "SAT0SUBNET":  "255.255.255.255",
                "ETH0IP" :  "10.61.34.38",
                "ETH0SUBNET": "255.255.255.252"
        },
        {
                "SVN" : 1013,
                "SAT0IP":  "100.81.7.236",
                "SAT0SUBNET":  "255.255.255.255",
                "ETH0IP" :  "10.61.50.38",
                "ETH0SUBNET": "255.255.255.252"
        },
        {
                "SVN" : 1011,
                "SAT0IP":  "100.72.3.28",
                "SAT0SUBNET":  "255.255.255.255",
                "ETH0IP" :  "192.168.1.6",
                "ETH0SUBNET": "255.255.255.0"
        }
        ]
  }`
  payload := strings.NewReader(payload_str)
  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("DID", "335577991")
  req.Header.Add("SAC", "193")
  req.Header.Add("Authorization", "Basic YWRtaW46YWRtaW4=")
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))


}

func main() {

	for {
		time.Sleep(3 * time.Second)
		send_inside_curl_command()
		time.Sleep(3 * time.Second)
		send_outside_curl_command()

	}	


}
