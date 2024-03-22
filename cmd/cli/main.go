package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/joho/godotenv/cmd/godotenv/internal/app/cli/config"
	"github.com/joho/godotenv/cmd/godotenv/pkg/env"
	"log"
)

func main() {
	fmt.Println("Hello world")

	config, err := config.InitConfig()

	if err != nil {
		log.Fatal("failed to init cli config", err)
	}

	fmt.Printf("%v", config)
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	fmt.Println(env.GetString("VARIABLE", ""))
}
