package usecase

import (
	"context"
	"fmt"

	"github.com/drybin/washington_changes_all/internal/domain/service"
	"github.com/drybin/washington_changes_all/pkg/wrap"
)

type IDayProcess interface {
	Process(ctx context.Context) error
}

type DayProcessUsecase struct {
	reportSenderService service.IReportSenderService
	service             service.IDayProcessService
}

func NewDayProcessUsecase(
	reportSenderService service.IReportSenderService,
	service service.IDayProcessService,
) *DayProcessUsecase {
	return &DayProcessUsecase{reportSenderService: reportSenderService, service: service}
}

func (u *DayProcessUsecase) Process(
	ctx context.Context,
) error {
	_, err := u.service.Process(ctx)
	if err != nil {
		return wrap.Errorf("failed to process day: %w", err)
	}

	//_, err := u.reportSenderService.Send(ctx, "First bot message")
	//if err != nil {
	//	return wrap.Errorf("failed to process: %w", err)
	//}

	fmt.Println("All done")
	return nil
}
