package main

import (
	"flag"

	"github.com/URL-shortening-service/internal/config"
)

func parseFlags(cfg *config.ConfigData) {

	flag.StringVar(&cfg.RunAddr, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&cfg.ShortAddr, "b", "http://localhost:8080", "base address of the resulting shortened URL")

	// парсим переданные серверу аргументы в зарегистрированные переменные
	flag.Parse()
}
