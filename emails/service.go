package emails

import "github.com/ahmedash95/go-channels/queue"

//EmailService ... email service
type EmailService struct {
	Queue chan queue.Queuable
}

//NewEmailService ... returns email service to send emails :D
func NewEmailService(q chan queue.Queuable) *EmailService {
	service := &EmailService{
		Queue: q,
	}

	return service
}

func (s EmailService) Send(e Email) {
	s.Queue <- e
}
