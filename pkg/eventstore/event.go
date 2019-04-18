package eventstore

type Event interface {
	GetAggregateType() string
	GetAggregateID() string
	GetEventNumber() int
	GetEventType() string
	GetData() ([]byte, error)
}

type BaseEvent struct {
	AggregateType string
	AggregateID   string
	EventNumber   int
}

func (e *BaseEvent) GetAggregateType() string { return e.AggregateType }
func (e *BaseEvent) GetAggregateID() string   { return e.AggregateID }
func (e *BaseEvent) GetEventNumber() int      { return e.EventNumber }
