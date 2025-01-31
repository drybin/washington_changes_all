package registry

import (
	"context"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/drybin/washington_changes_all/internal/adapter/pg"
	"github.com/drybin/washington_changes_all/internal/adapter/webapi"
	"github.com/drybin/washington_changes_all/internal/app/cli/config"
	"github.com/drybin/washington_changes_all/internal/app/cli/usecase"
	"github.com/drybin/washington_changes_all/internal/domain/service"
	"github.com/drybin/washington_changes_all/internal/domain/service/buy_strategy"
	"github.com/drybin/washington_changes_all/pkg/logger"
	"github.com/drybin/washington_changes_all/pkg/wrap"
	"github.com/go-resty/resty/v2"
	"github.com/jackc/pgx/v5"
)

type Container struct {
	Logger   logger.ILogger
	Usecases *Usecases
	Clean    func()
}

type Usecases struct {
	HelloWorld    *usecase.HelloWorld
	DayProcess    *usecase.DayProcessUsecase
	KucoinExplore *usecase.KucoinExploreUsecase
	GetBalance    *usecase.GetBalanceUsecase
}

func NewContainer(
	config *config.Config,
) (*Container, error) {
	log := logger.NewLogger()

	httpClient := resty.New()

	db, err := newDbConn(config)
	if err != nil {
		return nil, wrap.Errorf("failed to connect to Postgree: %w", err)
	}

	container := Container{
		Logger: log,
		Usecases: &Usecases{
			HelloWorld: usecase.NewHelloWorldUsecase(),
			DayProcess: usecase.NewDayProcessUsecase(
				service.NewReportSenderService(
					webapi.NewTelegramWebapi(
						httpClient,
						config.TgConfig.BotToken,
						config.TgConfig.ChatId,
					),
				),
				service.NewDayProcessService(
					pg.NewDaysRepository(db),
					pg.NewMarketsHistoryRepository(db),
					pg.NewCoinAvgPricesRepository(db),
					pg.NewBuyLogRepository(db),
					pg.NewCoinAmountRepository(db),
					webapi.NewKucoinWebapi(
						kucoin.NewApiService(
							kucoin.ApiKeyOption(config.KucoinConfig.Key),
							kucoin.ApiSecretOption(config.KucoinConfig.Secret),
							kucoin.ApiPassPhraseOption(config.KucoinConfig.Passphrase),
							kucoin.ApiKeyVersionOption(kucoin.ApiKeyVersionV2),
						),
					),
					buy_strategy.NewAmountEmptyStrategy(
						pg.NewCoinAvgPricesRepository(db),
						config.CoinConfig,
					),
					buy_strategy.NewMaxPriceDownStrategy(
						pg.NewCoinAvgPricesRepository(db),
						config.CoinConfig,
					),
					config.CoinConfig,
				),
			),
			KucoinExplore: usecase.NewKucoinExploreUsecase(
				service.NewCryptoExchangeServiceService(
					//stub.NewKucoinStub(),
					webapi.NewKucoinWebapi(
						kucoin.NewApiService(
							kucoin.ApiKeyOption(config.KucoinConfig.Key),
							kucoin.ApiSecretOption(config.KucoinConfig.Secret),
							kucoin.ApiPassPhraseOption(config.KucoinConfig.Passphrase),
							kucoin.ApiKeyVersionOption(kucoin.ApiKeyVersionV2),
						),
					),
				),
			),
			GetBalance: usecase.NewGetBalanceUsecase(
				service.NewCryptoExchangeServiceService(
					webapi.NewKucoinWebapi(
						kucoin.NewApiService(
							kucoin.ApiKeyOption(config.KucoinConfig.Key),
							kucoin.ApiSecretOption(config.KucoinConfig.Secret),
							kucoin.ApiPassPhraseOption(config.KucoinConfig.Passphrase),
							kucoin.ApiKeyVersionOption(kucoin.ApiKeyVersionV2),
						),
					),
				),
			),
		},
		Clean: func() {
		},
	}

	return &container, nil
}

func newDbConn(config *config.Config) (*pgx.Conn, error) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, config.PostgreeDsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
