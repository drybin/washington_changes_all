package command

import (
	"context"

	"github.com/drybin/washington_changes_all/internal/app/cli/usecase"
	"github.com/urfave/cli/v2"
)

func NewKucoinExploreCommand(uc usecase.IKucoinExplore) *cli.Command {
	return &cli.Command{
		Name:  "kucoin-explore",
		Usage: "kucoin-explore command",
		Flags: nil,
		Action: func(c *cli.Context) error {

			return uc.Process(
				context.Background(),
			)
		},
	}
}
