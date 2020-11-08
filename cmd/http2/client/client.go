package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"golang.org/x/net/http2"
	"io/ioutil"
	"net/http"
)

func main() {

	crt, _ := ioutil.ReadFile("./cmd/pem/public.crt")
	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(crt)

	tlsConfig := &tls.Config{
		RootCAs:            rootCAs,
		InsecureSkipVerify: false,
		ServerName:         "localhost",
	}

	h2Trans := &http2.Transport{
		TLSClientConfig:    tlsConfig,
		DisableCompression: true,
		AllowHTTP:          true,
	}

	client := &http.Client{
		Transport: h2Trans,
	}

	resp, err := client.Get("http://localhost:8080/test")
	if err != nil {
		fmt.Println(err)
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Printf("Code: %d\n", resp.StatusCode)
	fmt.Printf("Body: %s\n", body)

}
