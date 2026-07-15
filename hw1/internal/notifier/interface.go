package notifier

// Notifier - общий интерфейс для отправки уведомлений
type Notifier interface {
	Send(customer string) error
}
