package publisher

type Subscriber interface {
	Update(message string)
}

type EventSuscriber struct {
}

func (s *EventSuscriber) Update(message string) {
	println(message)
}
