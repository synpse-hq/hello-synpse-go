package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")

	r.GET("/", handleIndex)
	r.GET("/:name", handleIndex)
	r.Run(":8090")
}

func handleIndex(c *gin.Context) {
	name := c.Param("name")
	if name != "" {
		name = strings.TrimPrefix(c.Param("name"), "/")
	}
	c.HTML(http.StatusOK, "hellosynpse.html", gin.H{"Name": name})
}
