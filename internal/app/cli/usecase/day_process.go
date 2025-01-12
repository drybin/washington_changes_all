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
	strategyName := "(<i>применена стратегия макс падение от средней цены</i>)"
	strategyText := fmt.Sprintf(
		"монета упала на <b>%.2f%%</b> (<i>от средней цены %.2f</i>) текущая цена <b>%.2f</b>",
		day.CoinPriceChange,
		day.PrevCoinAvgPrice,
		day.Price,
	)
	if day.AmountEmptyStrategy {
		strategyName = fmt.Sprintf(
			"(<i>применена стратегия еще не покупалиs, монет которые еще не покупали в корзине %d</i>)",
			day.AmountEmptyStrategyCoinCount,
		)
		strategyText = fmt.Sprintf(
			"монета упала на <b>%.2f%%</b> (<i>от ath %.2f</i>) текущая цена <b>%.2f</b>",
			day.CoinPriceChange,
			day.CoinAth,
			day.Price,
		)
	}

	t := time.Now()
	return fmt.Sprintf("Сегодня <b>%d</b> день (<i>дата %d-%02d-%02d</i>)"+
		"(<i>баланс %.2f</i>)"+
		"Покупаем монету <b>%s</b> из корзины <b>%s</b>"+
		"%s"+
		"%s"+
		"Купили <b>%.2f</b>"+
		"Новая средняя цена монеты %.2f (<i>предыдущая %.2f</i>)"+
		"Новое количество монеты %.2f"+
		"На текущий момент куплено %d монет",
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
		day.CoinAmount,
		day.CoinCount,
	)
}
