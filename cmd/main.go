package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/golf1-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags)

	srv := server.New(logger)

	if err := srv.Server.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
