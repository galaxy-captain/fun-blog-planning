package main

import (
	"flag"
	"fmt"
	"fun-blog/internal/router"
)
import "github.com/gin-gonic/gin"

func main() {

	flag.Parse()

	server := gin.Default()
	err := router.Route(server)

	//err = server.Run(":8080")
	err = server.RunTLS(":8080", "./cmd/cert.pem", "./cmd/key.pem")
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

}
