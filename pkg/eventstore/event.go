package eventstore

type Event interface {
	GetAggregateID() string
	GetEventNumber() int
	GetEventType() string
	GetData() ([]byte, error)
}

type BaseEvent struct {
	AggregateID string
	EventNumber int
}

func (e *BaseEvent) GetAggregateID() string { return e.AggregateID }
func (e *BaseEvent) GetEventNumber() int    { return e.EventNumber }
