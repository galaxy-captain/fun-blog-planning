package main

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/http2"
	"io/ioutil"
	"net"
	"net/http"
)

func main() {

	h2Trans := &http2.Transport{
		DisableCompression: true,
		AllowHTTP:          true,
		DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
			return net.Dial(network, addr)
		},
	}

	client := &http.Client{
		Transport: h2Trans,
	}

	req, _ := http.NewRequest("POST", "http://localhost:8080/test", nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Printf("Code: %d\n", resp.StatusCode)
	fmt.Printf("Body:\n%s\n", body)

}
