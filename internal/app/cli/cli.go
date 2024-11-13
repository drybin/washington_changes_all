package cli

import (
	"github.com/drybin/washington_changes_all/internal/app/cli/config"
	"github.com/drybin/washington_changes_all/internal/app/cli/registry"
	"github.com/drybin/washington_changes_all/internal/presentation/command"
	"github.com/joho/godotenv"
	cliV2 "github.com/urfave/cli/v2"
	"log"
	"os"
)

const cliAppDesc = "cli tool for go-service-skeleton"

// example call go run --race ./cmd/cli/... users-sender --text ok-Luke
func Run(config *config.Config) error {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	cnt, err := registry.NewContainer(config)
	if err != nil {
		log.Fatal("failed to create cli container", err)
	}

	app := cliV2.NewApp()
	app.Name = config.ServiceName
	app.Usage = cliAppDesc
	app.Commands = []*cliV2.Command{
		command.NewUsersSenderCommand(cnt.Usecases.Users),
		command.NewDayProcessCommand(cnt.Usecases.DayProcess),
		command.NewKucoinExploreCommand(cnt.Usecases.KucoinExplore),
	}

	return app.Run(os.Args)
}
