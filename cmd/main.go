package main

import (
	"log"
	"os"

	"your_project/server"
)

func main() {
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags)

	srv := server.New(logger)

	if err := srv.Server.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
