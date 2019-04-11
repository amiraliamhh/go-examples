package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const endpoint = "http://127.0.0.1:8000/"

var body = map[string]interface{}{"name": "my name!"}

func sendGetReq() {
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

func sendPostReq() {
	reqBody, _ := json.Marshal(body)
	buff := bytes.NewBuffer([]byte(string(reqBody)))
	resp, err := http.Post(endpoint, "application/json", buff)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

func sendPutReq() {
	reqBody, _ := json.Marshal(body)
	buff := bytes.NewBuffer([]byte(string(reqBody)))
	req, err := http.NewRequest(http.MethodPut, endpoint, buff)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

func main() {
	sendGetReq()
	sendGetReq()
	sendPutReq()
}
