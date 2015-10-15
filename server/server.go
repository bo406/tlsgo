package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, This is an example of http service in golang!")
}

func main() {
	http.HandleFunc("/", handler)

	addr := ":4000"
	//	http.ListenAndServe(":3000", nil)
	log.Printf("start web server %s", addr)
	err := http.ListenAndServeTLS(addr, "server.crt", "server.key", nil)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
