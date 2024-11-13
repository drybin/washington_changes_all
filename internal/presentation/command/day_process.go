package command

import (
	"context"
	"github.com/drybin/washington_changes_all/internal/app/cli/usecase"
	"github.com/urfave/cli/v2"
)

func NewDayProcessCommand(uc usecase.IDayProcess) *cli.Command {
	return &cli.Command{
		Name:  "day-process",
		Usage: "day-process command",
		Flags: nil,
		Action: func(c *cli.Context) error {

			return uc.Process(
				context.Background(),
			)
		},
	}
}
