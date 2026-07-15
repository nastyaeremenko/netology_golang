package notifier

import "fmt"

// SMSSender отправляет уведомления по смс
type SMSSender struct{}

func (s *SMSSender) Send(customer string) error {
	fmt.Printf("SMS-уведомление отправлено клиенту %s\n", customer)
	return nil
}
