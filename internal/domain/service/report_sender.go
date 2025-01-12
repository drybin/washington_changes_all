package service

import (
	"github.com/drybin/washington_changes_all/internal/adapter/webapi"
	"github.com/drybin/washington_changes_all/pkg/wrap"
)

type ReportSenderService struct {
	telegramWebapi *webapi.TelegramWebapi
}

type IReportSenderService interface {
	Send(msg string) error
}

func NewReportSenderService(
	telegramWebapi *webapi.TelegramWebapi,
) IReportSenderService {
	return &ReportSenderService{telegramWebapi}
}

func (s *ReportSenderService) Send(msg string) error {
	_, err := s.telegramWebapi.Send(msg)

	if err != nil {
		return wrap.Errorf("failed to send report: %w", err)
	}

	return nil
}
