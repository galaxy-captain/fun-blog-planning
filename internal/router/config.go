package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandleRequest func(ctx *gin.Context)

func Route(server *gin.Engine) error {

	// load html files
	server.LoadHTMLGlob("template/*")

	api := server.Group("/api")
	api.GET("/blog/list", func(ctx *gin.Context) {
		//list := []string{"think-of-high-concurrency-request.md"}
		//ctx.JSON(http.StatusOK, list)
		ctx.HTML(http.StatusOK, "index.html", gin.H{"str": "hello world"})
	})

	return nil
}
