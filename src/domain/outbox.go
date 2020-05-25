package domain

type Outbox struct {
	ID            string
	EventType     string
	EventData     []byte
	AggregateID   string
	AggregateType string
}
