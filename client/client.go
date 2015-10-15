package main

import (
	"io/ioutil"
	"log"
	"net/http"
    "crypto/tls")

const URL = "https://localhost:4000"

func main() {
    tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

    client := &http.Client{Transport: tr}
	resp, err := client.Get(URL)
	if err != nil {
		log.Printf("get error: %v", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Printf(string(body))
}
