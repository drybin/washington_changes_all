package command

import (
	"context"

	"github.com/drybin/washington_changes_all/internal/app/cli/usecase"
	"github.com/drybin/washington_changes_all/internal/app/cli/usecase/dto"
	"github.com/urfave/cli/v2"
)

func NewUsersSenderCommand(sender usecase.IUserSender) *cli.Command {
	return &cli.Command{
		Name:  "users-sender",
		Usage: "users-sender command",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "text",
				Usage:    "text for write to console for example",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			text := c.String("text")

			return sender.Send(
				context.Background(),
				dto.SendUsersInput{
					Text: text,
				},
			)
		},
	}
}
