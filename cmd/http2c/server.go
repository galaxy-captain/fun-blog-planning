package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

func main() {

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(fmt.Sprintf("Header: %+v\n", r.Header)))
		_, _ = w.Write([]byte(fmt.Sprintf("Protocol: %s\n", r.Proto)))
	})

	h2s := &http2.Server{}
	server := &http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(handler, h2s),
	}

	log.Fatal(server.ListenAndServe())

}
