package site

import (
	"encoding/json"
	"fmt"

	"github.com/fewstera/divesites/pkg/eventstore"
)

type CreatedEvent struct {
	eventstore.BaseEvent
	createEventData
}
type createEventData struct {
	Name     string
	Location string
	Depth    float32
}

func (*CreatedEvent) GetEventType() string { return "SITE_CREATED" }
func (e *CreatedEvent) GetData() ([]byte, error) {
	d, err := json.Marshal(e.createEventData)
	if err != nil {
		return nil, fmt.Errorf("marshalling event data to json: %s", err)
	}

	return d, nil
}

type ReportAddedEvent struct {
	eventstore.BaseEvent
	reportAddedEventData
}
type reportAddedEventData struct {
	Reporter   string
	Visibility *float32
	Rating     int
	Notes      *string
}

func (*ReportAddedEvent) GetEventType() string { return "SITE_REPORT_ADDED" }
func (e *ReportAddedEvent) GetData() ([]byte, error) {
	d, err := json.Marshal(e.reportAddedEventData)
	if err != nil {
		return nil, fmt.Errorf("marshalling event data to json: %s", err)
	}

	return d, nil
}
