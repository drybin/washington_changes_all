package service

import (
	"github.com/drybin/washington_changes_all/internal/adapter/webapi"
	"github.com/drybin/washington_changes_all/pkg/wrap"
)

type ReportSenderService struct {
	telegramWebapi *webapi.TelegramWebapi
}

type IReportSenderService interface {
	Send(msg string) (bool, error)
}

func NewReportSenderService(
	telegramWebapi *webapi.TelegramWebapi,
) IReportSenderService {
	return &ReportSenderService{telegramWebapi}
}

func (s *ReportSenderService) Send(msg string) (bool, error) {
	_, err := s.telegramWebapi.Send(msg)

	if err != nil {
		return false, wrap.Errorf("failed to send report: %w", err)
	}

	return true, nil
}
