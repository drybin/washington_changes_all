package command

import (
	"context"

	"github.com/drybin/washington_changes_all/internal/app/cli/usecase"
	"github.com/urfave/cli/v2"
)

func NewDayProcessSellCommand(uc usecase.IDayProcessSell) *cli.Command {
	return &cli.Command{
		Name:  "day-process-sell",
		Usage: "day-process-sell command",
		Flags: nil,
		Action: func(c *cli.Context) error {

			return uc.Process(
				context.Background(),
			)
		},
	}
}
