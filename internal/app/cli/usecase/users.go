package usecase

import (
	"context"
	"github.com/drybin/washington_changes_all/internal/app/cli/usecase/dto"
	"log"
)

type IUserSender interface {
	Send(ctx context.Context, input dto.SendUsersInput) error
}

type UsersUsecase struct {
}

func NewUsersUsecase() *UsersUsecase {
	return &UsersUsecase{}
}

func (u *UsersUsecase) Send(
	_ context.Context,
	input dto.SendUsersInput,
) error {
	log.Println(input.Text)

	return nil
}
