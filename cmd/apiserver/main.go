package main

import (
	"fmt"

	"http-rest-api-go/internal/config"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Todo: init config
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// Todo: init Logger

	// Todo: init storage

	// Tode: init router

	// Todo: run server

}
