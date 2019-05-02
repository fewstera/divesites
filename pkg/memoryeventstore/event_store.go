package memoryeventstore

import (
	"errors"
	"fmt"
	"sync"

	"github.com/fewstera/divesites/pkg/eventstore"
)

type EventStore struct {
	sync.RWMutex
	eventsByAggregateType eventsByAggregateTypeMap
	transactionBackupMap  eventsByAggregateTypeMap
}

type eventsByAggregateTypeMap map[string]eventsByAggregateIdMap
type eventsByAggregateIdMap map[string][]eventstore.Event

func New() *EventStore {
	es := &EventStore{}
	es.eventsByAggregateType = make(eventsByAggregateTypeMap)
	return es
}

func (es *EventStore) Store(events []eventstore.Event) error {
	es.Lock()
	defer es.Unlock()

	es.beginTransaction()
	defer es.endTransaction()

	for i, event := range events {
		byID, ok := es.eventsByAggregateType[event.GetAggregateType()]
		if !ok {
			byID = make(eventsByAggregateIdMap)
		}
		aggregateEvents := byID[event.GetAggregateID()]

		lastCommitEventNumber := 0
		if len(aggregateEvents) > 0 {
			lastCommitEventNumber = aggregateEvents[len(aggregateEvents)-1].GetEventNumber()
		}
		if event.GetEventNumber() != lastCommitEventNumber+1 {
			es.rollbackTransaction()
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

func (es *EventStore) beginTransaction() error {
	if es.transactionBackupMap != nil {
		return errors.New("tried to start a new transaction but one already exists")
	}

	es.transactionBackupMap = make(eventsByAggregateTypeMap)
	copyEventsByAggregateTypeMap(es.transactionBackupMap, es.eventsByAggregateType)

	return nil
}

func (es *EventStore) endTransaction() {
	es.transactionBackupMap = nil
}

func (es *EventStore) rollbackTransaction() {
	es.eventsByAggregateType = make(eventsByAggregateTypeMap)
	copyEventsByAggregateTypeMap(es.eventsByAggregateType, es.transactionBackupMap)
}

func copyEventsByAggregateTypeMap(to, from eventsByAggregateTypeMap) {
	for aggregateType, aggregateIDMap := range from {
		newAggregateIDMap := make(eventsByAggregateIdMap)
		to[aggregateType] = newAggregateIDMap

		for aggregateID, events := range aggregateIDMap {
			newAggregateIDMap[aggregateID] = make([]eventstore.Event, len(events))
			copy(newAggregateIDMap[aggregateID], events)
		}
	}
}
