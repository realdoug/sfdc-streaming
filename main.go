package main

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

func main() {
	url := "http://localhost:9292/faye"
	var params = `{"channel":"/meta/handshake", "supportedConnectionTypes":["long-polling"], "version":"1.0"}`
	request := gorequest.New()
	_, body, _ := request.Post(url).Set("Content-Type", "application/json").Send(params).End()
	var data []map[string]interface{}
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		panic(err)
	}
	fmt.Println(data)
}
