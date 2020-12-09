package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
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

	h1Trans := &http.Transport{
		TLSClientConfig:    tlsConfig,
		DisableCompression: true,
	}

	client := &http.Client{
		Transport: h1Trans,
	}

	resp, err := client.Get("https://localhost:8443/test")
	if err != nil {
		fmt.Println(err)
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Printf("Code: %d\n", resp.StatusCode)
	fmt.Printf("Body: %s\n", body)

}
