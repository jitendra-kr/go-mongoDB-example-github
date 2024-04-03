package main

import (
	"log"

	"github.com/go-mongoDB-example-github/internal"
)

var server *internal.Server

func init() {
	server = internal.New()
}

func main() {
	if err := server.Start(); err != nil {
		log.Fatal("service failure", err)
	}
}
