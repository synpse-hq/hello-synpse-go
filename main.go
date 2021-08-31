package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func run() error {
	certFile, certSet := os.LookupEnv("TLS_CRT")
	keyFile, keySet := os.LookupEnv("TLS_KEY")

	port := ":8080"
	if val, ok := os.LookupEnv("PORT"); ok {
		port = val
	}

	r := gin.Default()

	r.LoadHTMLGlob("./templates/*")

	r.GET("/", handleIndex)
	r.GET("/:name", handleIndex)
	if !certSet || !keySet {
		log.Print("TLS_CRT and TLS_KEY  not set, running in insecure mode")
		r.Run(port)
	} else {
		r.RunTLS(port, certFile, keyFile)
	}
	return nil
}

func handleIndex(c *gin.Context) {
	name := c.Param("name")
	if name != "" {
		name = strings.TrimPrefix(c.Param("name"), "/")
	}
	c.HTML(http.StatusOK, "hellosynpse.html", gin.H{"Name": name})
}
