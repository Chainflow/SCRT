package alerting

// TelegramAlerter is an interface definition for telegram related actions
// like Send telegram alert
type TelegramAlerter interface {
	Send(msgText, botToken string, alerthnatID int64) error
}

type telegramAlert struct{}

// NewTelegramAlerter returns a new instance for telegramAlert
func NewTelegramAlerter() *telegramAlert {
	return &telegramAlert{}
}
