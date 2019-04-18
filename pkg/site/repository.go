package site

import (
	"fmt"

	"github.com/fewstera/divesites/pkg/eventstore"
)

type Repository struct {
	es eventstore.EventStore
}

func NewRepository(es eventstore.EventStore) *Repository {
	return &Repository{
		es: es,
	}
}

func (r *Repository) Get(aggregateID string) (*Site, error) {
	events, err := r.es.AggregateEvents(Aggregate, aggregateID)
	if err != nil {
		return nil, fmt.Errorf("getting events for aggregate: %s", err)
	}

	s := &Site{}
	s.Apply(events)

	return s, nil
}
