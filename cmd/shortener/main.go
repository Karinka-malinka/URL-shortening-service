package main

import (
	"github.com/URL-shortening-service/internal/handlers"
)

func main() {

	if err := handlers.Run(); err != nil {
		panic(err)
	}

}
