package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	crt, _ := ioutil.ReadFile("./cmd/pem/public.crt")
	key, _ := ioutil.ReadFile("./cmd/pem/private.key")
	cert, _ := tls.X509KeyPair(crt, key)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
	}

	server := &http.Server{
		TLSConfig: tlsConfig,
	}

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(fmt.Sprintf("Protocol: %s", r.Proto)))
	})

	server.ListenAndServeTLS("", "")

}
