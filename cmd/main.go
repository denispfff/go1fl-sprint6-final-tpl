package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	mainLogger := log.New(
		os.Stdout,
		"server: ",
		log.LstdFlags,
	)
	srv := server.New(mainLogger)

	if err := srv.HttpServer.ListenAndServe(); err != nil {
		mainLogger.Fatal(err)
	}
}
