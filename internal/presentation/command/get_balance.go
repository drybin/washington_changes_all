package command

import (
	"context"

	"github.com/drybin/washington_changes_all/internal/app/cli/usecase"
	"github.com/urfave/cli/v2"
)

func NewGetBalanceCommand(u usecase.IGetBalance) *cli.Command {
	return &cli.Command{
		Name:  "get-balance",
		Usage: "get-balance command",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			return u.Get(context.Background())
		},
	}
}
