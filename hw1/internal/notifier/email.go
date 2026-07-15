package notifier

import "fmt"

// EmailSender отправляет уведомления по почте
type EmailSender struct{}

func (s *EmailSender) Send(customer string) error {
	fmt.Printf("Email-уведомление отправлено клиенту %s\n", customer)
	return nil
}
