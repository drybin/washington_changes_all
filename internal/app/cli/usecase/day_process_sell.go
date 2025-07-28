package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/drybin/washington_changes_all/internal/domain/model"
	"github.com/drybin/washington_changes_all/internal/domain/service"
	"github.com/drybin/washington_changes_all/pkg/wrap"
)

type IDayProcessSell interface {
	Process(ctx context.Context) error
}

type DayProcessSellUsecase struct {
	reportSenderService service.IReportSenderService
	service             service.IDayProcessSellService
}

func NewDayProcessSellUsecase(
	reportSenderService service.IReportSenderService,
	service service.IDayProcessSellService,
) *DayProcessSellUsecase {
	return &DayProcessSellUsecase{reportSenderService: reportSenderService, service: service}
}

func (u *DayProcessSellUsecase) Process(
	ctx context.Context,
) error {
	daySellResult, err := u.service.Process(ctx)
	if err != nil {
		return wrap.Errorf("failed to process day: %w", err)
	}

	if len(daySellResult.Coins) > 0 {
		msg := generateReportSell(*daySellResult)
		err = u.reportSenderService.Send(msg)
		if err != nil {
			return wrap.Errorf("failed to process: %w", err)
		}
	}

	fmt.Println("All done")
	return nil
}

func generateReportSell(day model.DaySellResult) string {
	newLine := "\n"
	sellReport := "<b>Продали монеты:</b>" + newLine

	for _, coinInfo := range day.Coins {
		sellReport += fmt.Sprintf(
			"монета: %s"+newLine+
				"количество %.6f(<i>предыдущее %.6f</i>)"+newLine+
				"текущая цена %.6f(<i>средняя %.6f</i>)"+newLine+
				"продали на %.6f",
			coinInfo.CoinName,
			coinInfo.CoinNewAmount,
			coinInfo.CoinPrevAmount,
			coinInfo.CoinCurrentPrice,
			coinInfo.CoinAvgPrice,
			coinInfo.SellUsdtAmount,
		)
	}

	t := time.Now()
	return fmt.Sprintf("Сегодня <b>%d</b> день (<i>дата %d-%02d-%02d</i>)"+newLine+
		"(<i>баланс %.2f</i>)"+newLine+
		"%s",
		day.DayInfo.ID,
		t.Year(),
		t.Month(),
		t.Day(),
		day.Balance,
		sellReport,
	)
}
