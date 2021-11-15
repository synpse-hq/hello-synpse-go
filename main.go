package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func run() error {
	// override port if provided
	port := ":8080"
	if val, ok := os.LookupEnv("PORT"); ok {
		port = val
	}

	// create filesystem router
	fs := http.FileServer(http.Dir("./ui"))
	http.Handle("/", fs)

	// read certificate variables and if set - run in https mode.
	// else - run in plain http mode.
	certFile, certSet := os.LookupEnv("TLS_CRT")
	keyFile, keySet := os.LookupEnv("TLS_KEY")
	if certSet || keySet {
		log.Println("Listening TLS on " + port)
		http.ListenAndServeTLS(port, certFile, keyFile, nil)
	} else {
		log.Println("Listening on " + port)
		err := http.ListenAndServe(port, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
