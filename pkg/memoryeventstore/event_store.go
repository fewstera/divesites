package memoryeventstore

import "github.com/fewstera/divesites/pkg/eventstore"

//type EventStore interface {
//Store(event Event) error
//AggregateEvents(aggregateType, aggregateID string) []Event
//}

type EventStore struct {
	eventsByAggregateType map[string]eventsByAggregateId
}

type eventsByAggregateId map[string][]eventstore.Event

func New() *EventStore {
	es := &EventStore{}
	es.eventsByAggregateType = make(map[string]eventsByAggregateId)
	return es
}

func (es *EventStore) Store() error {

}
