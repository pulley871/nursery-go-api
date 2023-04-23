package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("Bye")
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSONP(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
