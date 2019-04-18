package memoryeventstore

// TODO:
//  - Make Store() transactional - If sent 5 events, but 3 event fails to store, then all
//    events 1 and 2 should not be commited to the event store
import (
	"fmt"
	"sync"

	"github.com/fewstera/divesites/pkg/eventstore"
)

type EventStore struct {
	sync.RWMutex
	eventsByAggregateType map[string]eventsByAggregateId
}

type eventsByAggregateId map[string][]eventstore.Event

func New() *EventStore {
	es := &EventStore{}
	es.eventsByAggregateType = make(map[string]eventsByAggregateId)
	return es
}

func (es *EventStore) Store(events []eventstore.Event) error {
	es.Lock()
	defer es.Unlock()

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
	es.RLock()
	defer es.RUnlock()

	byAggregateId, ok := es.eventsByAggregateType[aggregateType]
	if !ok {
		return nil, fmt.Errorf("no events for aggregate type '%s' found", aggregateType)
	}

	aggregateEvents, ok := byAggregateId[aggregateID]
	if !ok {
		return nil, fmt.Errorf("no events for aggregate id '%s' found", aggregateID)
	}

	return aggregateEvents, nil
}
