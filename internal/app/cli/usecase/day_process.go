package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/drybin/washington_changes_all/internal/domain/model"
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
	dayResult, err := u.service.Process(ctx)
	if err != nil {
		return wrap.Errorf("failed to process day: %w", err)
	}

	msg := generateReport(*dayResult)
	err = u.reportSenderService.Send(msg)
	if err != nil {
		return wrap.Errorf("failed to process: %w", err)
	}

	fmt.Println("All done")
	return nil
}

func generateReport(day model.DayResult) string {
	newLine := "\n"
	strategyName := "(<i>применена стратегия макс падение от средней цены</i>)" + newLine
	strategyText := fmt.Sprintf(
		"монета упала на <b>%.2f%%</b> (<i>от средней цены %.4f</i>) текущая цена <b>%.4f</b>"+newLine,
		day.CoinPriceChange,
		day.PrevCoinAvgPrice,
		day.Price,
	)
	if day.AmountEmptyStrategy {
		strategyName = fmt.Sprintf(
			"(<i>применена стратегия еще не покупали, монет которые еще не покупали в корзине %d</i>)"+newLine,
			day.AmountEmptyStrategyCoinCount,
		)
		strategyText = fmt.Sprintf(
			"монета упала на <b>%.2f%%</b> (<i>от ath %.4f</i>) текущая цена <b>%.4f</b>"+newLine,
			day.CoinPriceChange,
			day.CoinAth,
			day.Price,
		)
	}

	t := time.Now()
	return fmt.Sprintf("Сегодня <b>%d</b> день (<i>дата %d-%02d-%02d</i>)"+newLine+
		"(<i>баланс %.2f</i>)"+newLine+
		"Покупаем монету <b>%s</b> из корзины <b>%s</b>"+newLine+
		"%s"+
		"%s"+
		"Купили <b>%.4f</b>"+newLine+
		"Новая средняя цена монеты %.4f (<i>предыдущая %.4f</i>)"+newLine+
		"Новое количество монеты %.7f"+newLine+
		"На текущий момент куплено %d монет"+newLine,
		day.DayNumber,
		t.Year(),
		t.Month(),
		t.Day(),
		day.Balance,
		day.CoinName.String(),
		day.Tier.String(),
		strategyName,
		strategyText,
		day.Amount,
		day.CoinAvgPrice,
		day.PrevCoinAvgPrice,
		day.Amount,
		day.CoinCount,
	)
}
