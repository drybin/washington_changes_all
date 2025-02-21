package usecase

import (
    "context"
    "log"
    
    "github.com/drybin/washington_changes_all/internal/domain/repo"
    "github.com/drybin/washington_changes_all/pkg/wrap"
)

type IGetBalance interface {
    Get(ctx context.Context) error
}

type GetBalanceUsecase struct {
    repo repo.ICryptoExchangeRepository
}

func NewGetBalanceUsecase(repo repo.ICryptoExchangeRepository) *GetBalanceUsecase {
    return &GetBalanceUsecase{
        repo: repo,
    }
}

func (u *GetBalanceUsecase) Get(ctx context.Context) error {
    log.Println("Получаем баланс")
    
    balance, err := u.repo.GetBalance(ctx)
    if err != nil {
        return wrap.Errorf("failed to get balance: %w", err)
    }
    
    log.Printf("баланс: %.2f$", balance)
    return nil
}
