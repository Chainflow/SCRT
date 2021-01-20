package server

import (
	"log"

	"github.com/Chainflow/SCRT/alerting"
	"github.com/Chainflow/SCRT/config"
)

// SendTelegramAlert sends the alert to telegram chat
func SendTelegramAlert(msg string, cfg *config.Config) error {
	if err := alerting.NewTelegramAlerter().Send(msg, cfg.Telegram.BotToken, cfg.Telegram.ChatID); err != nil {
		log.Printf("failed to send telegram alert: %v", err)
		return err
	}
	return nil
}
