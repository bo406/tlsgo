package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const ADDR = ":4000"

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "第一个TLS Server, Oh Yeah!")
}

func main() {

	pool := x509.NewCertPool()
	caCertPath := "ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Printf("ReadFlie ca.crt error: %v", err)
		return
	}

	pool.AppendCertsFromPEM(caCrt)

	addr := ":4000"
	s := &http.Server{
		Addr:    ADDR,
		Handler: &myhandler{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	log.Printf("start https web server %s", addr)
	err = s.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		log.Fatalf("%v", err)
	}
}
