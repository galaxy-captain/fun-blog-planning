package main

import (
	"flag"
	"fmt"
	"net/http"
)
import "github.com/gin-gonic/gin"

func main() {

	flag.Parse()

	server := gin.New()
	api := server.Group("/api")
	api.GET("/blog/list", func(ctx *gin.Context) {
		list := []string{"think-of-high-concurrency-request.md"}
		ctx.JSON(http.StatusOK, list)
	})

	err := server.Run(":8080")
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

}
