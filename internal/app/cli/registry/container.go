package registry

import (
	"github.com/drybin/washington_changes_all/internal/app/cli/config"
	"github.com/drybin/washington_changes_all/internal/app/cli/usecase"
	"github.com/drybin/washington_changes_all/pkg/logger"
)

type Container struct {
	Logger   logger.ILogger
	Usecases *Usecases
	Clean    func()
}

type Usecases struct {
	Users *usecase.UsersUsecase
}

func NewContainer(
	config *config.Config,
) (*Container, error) {
	log := logger.NewLogger()

	container := Container{
		Logger: log,
		Usecases: &Usecases{
			Users: usecase.NewUsersUsecase(),
		},
		Clean: func() {
		},
	}

	return &container, nil
}
