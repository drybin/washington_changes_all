package repo

import (
	"context"

	"github.com/drybin/washington_changes_all/internal/domain/model"
)

type IDaysRepository interface {
	GetLastDayInfo(ctx context.Context) (*model.Day, error)
	Save(ctx context.Context, day model.Day) (*model.Day, error)
}
