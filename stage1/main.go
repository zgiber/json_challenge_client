package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	server = "127.0.0.1:4000"
)

func main() {

	requestUrl := url.URL{}
	requestUrl.Host = server
	requestUrl.Scheme = "http"
	requestUrl.Path = "/stage1/data.json"

	stage1Data, err := http.Get(requestUrl.String())
	if err != nil {
		log.Println("%s - Sorry about that", err.Error())
		return
	}

	// YOUR CODE HERE
	_ = stage1Data //

	responseData := []byte{} // FOR YOUR RESPONSE
	body := bytes.NewBuffer(responseData)

	responseUrl := requestUrl
	responseUrl.Path = "/stage1/submit.json"
	stage1SubmitResponse, err := http.Post(responseUrl.String(), "application/json", body)
	if err != nil {
		log.Println("%s - Sorry about that", err.Error())
		return
	}

	responseData, _ = ioutil.ReadAll(stage1SubmitResponse.Body)
	log.Println(string(responseData))
}
