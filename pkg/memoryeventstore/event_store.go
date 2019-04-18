package memoryeventstore

import (
	"fmt"

	"github.com/fewstera/divesites/pkg/eventstore"
)

//type EventStore interface {
//Store(events []Event) error
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

func (es *EventStore) Store(events []eventstore.Event) error {
	for i, event := range events {
		byID, ok := es.eventsByAggregateType[event.GetAggregateType()]
		if !ok {
			byID = make(eventsByAggregateId)
		}
		aggregateEvents := byID[event.GetAggregateID()]

		lastCommitEventNumber := 0
		if len(aggregateEvents) > 0 {
			lastCommitEventNumber = aggregateEvents[len(aggregateEvents)-1].GetEventNumber()
		}
		if event.GetEventNumber() != lastCommitEventNumber+1 {
			return fmt.Errorf(
				"event at index %d: invalid event number: expected %d, got %d",
				i, lastCommitEventNumber+1, event.GetEventNumber(),
			)
		}

		aggregateEvents = append(aggregateEvents, event)
		byID[event.GetAggregateID()] = aggregateEvents
		es.eventsByAggregateType[event.GetAggregateType()] = byID
	}

	return nil
}

func (es *EventStore) AggregateEvents(aggregateType, aggregateID string) ([]eventstore.Event, error) {
	return es.eventsByAggregateType[aggregateType][aggregateID], nil
}
