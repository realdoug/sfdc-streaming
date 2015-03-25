package main

import (
	"fmt"
	"github.com/realdoug/go-force/force"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type SalesforceCreds struct {
	ClientId      string `yaml:"client_id"`
	ClientSecret  string `yaml:"client_secret"`
	Username      string
	Password      string
	SecurityToken string `yaml:"security_token"`
}

func main() {
	c := make(chan int)
	forceApi := connectToSalesforce()
	forceApi.ConnectToStreamingApi()

	callback := func(data []byte, args ...interface{}) {
		fmt.Println(string(data))
	}

	forceApi.SubscribeToPushTopic("Tasks", callback)
	<-c
}

func connectToSalesforce() *force.ForceApi {
	sfyamlfile, _ := ioutil.ReadFile("salesforce_creds.yml")
	var creds = SalesforceCreds{}
	yaml.Unmarshal(sfyamlfile, &creds)
	forceApi, err := force.Create(
		"v33.0",
		creds.ClientId,
		creds.ClientSecret,
		creds.Username,
		creds.Password,
		creds.SecurityToken,
		"production",
	)
	if err != nil {
		panic(err)
	}

	return forceApi
}
