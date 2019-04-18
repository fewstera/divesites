package site

import (
	"github.com/fewstera/divesites/pkg/eventstore"
)

type Site struct {
	ID       string
	Name     string
	Location string
	Depth    float32

	changes []eventstore.Event
}

var Aggregate = "site"

func New(id, name, location string, depth float32) *Site {
	s := &Site{}

	createdEvent := &CreatedEvent{}
	createdEvent.AggregateID = id
	createdEvent.AggregateType = Aggregate
	createdEvent.EventNumber = 1
	createdEvent.Name = name
	createdEvent.Location = location
	createdEvent.Depth = depth

	s.applyEvents(true, []eventstore.Event{createdEvent})

	return s
}

func (s *Site) Apply(events []eventstore.Event) {
	s.applyEvents(false, events)
}

func (s *Site) applyEvents(isNew bool, events []eventstore.Event) {
	for _, event := range events {
		switch e := event.(type) {
		case *CreatedEvent:
			s.ID = e.AggregateID
			s.Name = e.Name
			s.Location = e.Location
			s.Depth = e.Depth
		}
	}

	if isNew {
		s.changes = append(s.changes, events...)
	}
}

func (s *Site) UncommitedChanges() []eventstore.Event {
	return s.changes
}

func (s *Site) ClearChanges() {
	s.changes = nil
}
