type EventStore interface {
	Store(event Event) error
	AggregateEvents(aggregateType, aggregateID string) []Event
}
