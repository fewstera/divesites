package site

import (
	"github.com/fewstera/divesites/pkg/eventstore"
)

type Site struct {
	ID       string   `json:"id"`
	Version  int      `json:"version"`
	Name     string   `json:"name"`
	Location string   `json:"location"`
	Depth    float32  `json:"depth"`
	Reports  []Report `json:"reports"`

	changes []eventstore.Event
}

type Report struct {
	Reporter   string
	Visibility *float32
	Rating     int
	Notes      *string
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

func (s *Site) AddReport(report Report) {
	reportAddedEvent := &ReportAddedEvent{}
	reportAddedEvent.AggregateID = s.ID
	reportAddedEvent.AggregateType = Aggregate
	reportAddedEvent.EventNumber = s.Version + 1
	reportAddedEvent.Reporter = report.Reporter
	reportAddedEvent.Visibility = report.Visibility
	reportAddedEvent.Rating = report.Rating
	reportAddedEvent.Notes = report.Notes

	s.applyEvents(true, []eventstore.Event{reportAddedEvent})
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
			s.Reports = []Report{}
		case *ReportAddedEvent:
			report := Report{
				Reporter:   e.Reporter,
				Visibility: e.Visibility,
				Rating:     e.Rating,
				Notes:      e.Notes,
			}
			s.Reports = append(s.Reports, report)
		}

		s.Version = event.GetEventNumber()
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
