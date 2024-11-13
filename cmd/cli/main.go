package main

import (
	"github.com/drybin/washington_changes_all/internal/app/cli"
	"github.com/drybin/washington_changes_all/internal/app/cli/config"
	"github.com/joho/godotenv"
	"log"
)

const (
	appName = "go-service-skeleton-cli"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load env", err)
	}

	config, err := config.InitConfig()
	if err != nil {
		log.Fatal("failed to init cli config", err)
	}

	if err := cli.Run(config); err != nil {
		log.Fatal("failed to run cli app: ", err)
	}
}

//package main
//
//import (
//	"fmt"
//	"github.com/drybin/washington_changes_all/internal/app/cli/config"
//	"github.com/drybin/washington_changes_all/pkg/env"
//	"github.com/joho/godotenv"
//	"log"
//)
//
//func main() {
//	fmt.Println("Hello world")
//
//	if err := godotenv.Load(); err != nil {
//		log.Println(err)
//	}
//
//	config, err := config.InitConfig()
//
//	if err != nil {
//		log.Fatal("failed to init cli config", err)
//	}
//
//	fmt.Printf("%v", config)
//	fmt.Println(env.GetString("VARIABLE", ""))
//}
