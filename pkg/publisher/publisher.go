package publisher

type Publisher interface {
	Publish() error
}

// EventPublisher is a simple interface for publishing messages
type EventPublisher struct {
	message    chan string
	suscribers []Subscriber
}

func (p *EventPublisher) Suscribe(s Subscriber) error {
	p.suscribers = append(p.suscribers, s)
	return nil
}

func (p *EventPublisher) Publish() error {
	message := <-p.message

	for _, subscriber := range p.suscribers {
		subscriber.Update(message)
	}

	return nil
}

func NewPublisher(publish func(string) error) Publisher {
	e := EventPublisher{}
	return &e
}
