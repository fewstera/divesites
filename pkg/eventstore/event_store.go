package eventstore

type EventStore interface {
	Store(events []Event) error
	AggregateEvents(aggregateType, aggregateID string) ([]Event, error)
}
